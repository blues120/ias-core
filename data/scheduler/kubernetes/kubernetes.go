package kubernetes

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler/agent"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/convert"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/k8s"
	"gopkg.in/yaml.v3"
	k8sCoreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	redisKeyJobCache = "ias:stoppingjobs" // 缓存停止中的任务用于状态显示，Pod 无此状态

	TaskPodLabel         = "ias-task"
	TaskPodLabelSelector = "app=ias-task"
)

type scheduler struct {
	bc *conf.Bootstrap

	client *kubernetes.Clientset

	rdb *redis.Client

	log *log.Helper
}

func newKubernetes(kubeCfg *conf.Data_Kubernetes) *kubernetes.Clientset {
	// 指定配置文件
	config, err := clientcmd.BuildConfigFromFlags("", kubeCfg.KubeConfig)
	if err != nil {
		panic(err.Error())
	}
	// 创建 client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func newRedisClient(c *conf.Data_Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Password:     c.Password,
		DB:           int(c.Db),
		WriteTimeout: c.WriteTimeout.AsDuration(),
		ReadTimeout:  c.ReadTimeout.AsDuration(),
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}

func NewScheduler(bc *conf.Bootstrap, rdb *redis.Client, logger log.Logger) biz.SchedulerRepo {
	if bc == nil || bc.Data == nil || logger == nil || bc.Data.Kubernetes == nil || bc.Data.Redis == nil {
		return nil
	}

	return &scheduler{
		bc:     bc,
		client: newKubernetes(bc.Data.Kubernetes),
		rdb:    rdb,
		log:    log.NewHelper(logger),
	}
}

// Close 关闭
func (r *scheduler) Close() error {
	if err := r.rdb.Close(); err != nil {
		r.log.Errorf("failed to close redis: %s, redis info: %v", err, r.bc.Data.Redis)
		return err
	}
	return nil
}

// Start 启动任务
func (r *scheduler) Start(ctx context.Context, ta *biz.Task) error {
	config := agent.TaskToAgentConfig(r.bc, ta, nil)
	switch ta.Type {
	case biz.TaskTypeFrame:
		return r.startFrameTask(ctx, ta)
	case biz.TaskTypeStream:
		return r.startStreamTask(ctx, ta, config)
	default:
		return fmt.Errorf("unsupported task type: %v", ta.Type)
	}
}

// 摄像头监控区域
type TargetAreas struct {
	TargetArea [][]struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
	ExcludeArea [][]struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	}
}

// 获取摄像头监控区域坐标
func getCameraImgBox(ca biz.TaskCamera) string {
	return ca.MultiImgBox
}

// 根据任务id，返回job名称
func getFrameJobName(jobId uint) string {
	return fmt.Sprintf("%d", jobId)
}

type FrameJobParam struct {
	Task           *biz.Task
	Env            map[string]string
	FfmpegCommand  []string
	SideCarCommand []string
	InitCommand    []string
	CameraCount    int
}

type cameraAttr struct {
	Id       int64  `json:"id"` // 摄像机ID
	Rate     string `json:"rate"`
	Source   string `json:"source"`
	Codec    string `json:"codec"`    // 摄像机编码格式
	GpuCodec string `json:"gpuCodec"` // 摄像机编码格式
	Width    int    `json:"width"`    // 像素宽
	Height   int    `json:"height"`   // 像素高
}

// 生成run ffmpeg 需要的camera list
func (r *scheduler) genCameraConf(ta *biz.Task, interval float64) []byte {
	var ffmpegCameraArgs []*cameraAttr

	for _, c := range ta.Cameras {
		ffmpegCameraArgs = append(ffmpegCameraArgs, &cameraAttr{
			Id:       int64(c.Id),
			Rate:     fmt.Sprintf("%.3f", interval),
			Source:   c.Camera.StreamingProtocol.Source(),
			Codec:    c.Camera.StreamingInfo.CodecName,
			GpuCodec: fmt.Sprintf("%s_cuvid", c.Camera.StreamingInfo.CodecName),
			Height:   int(c.Camera.StreamingInfo.Height),
			Width:    int(c.Camera.StreamingInfo.Width),
		})
	}

	cameraConf, _ := json.Marshal(ffmpegCameraArgs)
	return cameraConf
}

type algoExtra struct {
	FaceGroupId   string `json:"face_group_id"`   // 人脸底库 id，逗号分隔支持多底库
	FaceGroupName string `json:"face_group_name"` // 人脸底库名，逗号分隔支持多底库
}

func (r *scheduler) startFrameTask(ctx context.Context, ta *biz.Task) error {
	interval := convert.ConvertDropFrameIntervalToFPS(int(ta.Algo.Interval))

	// 生成摄像头坐标配置
	cameraBoxList := make(map[uint]string, 0)
	cameraImageDir := make([]string, 0)
	for _, v := range ta.Cameras {
		if v.MultiImgBox != "" {
			boxes := getCameraImgBox(v)
			if boxes != "" {
				cameraBoxList[uint(v.Id)] = boxes
			}
		}
		cameraImageDir = append(cameraImageDir, fmt.Sprint(r.bc.Data.Kubernetes.FrameTaskBaseDir, v.Id))
	}
	boxJson, _ := jsoniter.Marshal(cameraBoxList)

	// 生成摄像头Attr配置
	cameraCodecList := make(map[uint]string, 0)
	for _, v := range ta.Cameras {
		cameraCodecList[uint(v.Id)] = v.Camera.StreamingInfo.CodecName
	}

	// 获取人脸库配置
	var extra algoExtra
	if ta.Algo.Extra != "" {
		err := json.Unmarshal([]byte(ta.Algo.Extra), &extra)
		if err != nil {
			r.log.Errorf("failed to unmarshal algo extra: %s", err)
		}
	}

	env := make(map[string]string)
	awsS3 := r.bc.Data.Oss.GetAwsS3()
	if awsS3 != nil {
		env["Endpoint"] = awsS3.Endpoint
		env["AccessKey"] = awsS3.Ak
		env["SecretAccessKey"] = awsS3.Sk
		env["Bucket"] = awsS3.Bucket
		env["TrainDataBucket"] = awsS3.TrainDataBucket
	}
	env["job_id"] = getFrameJobName(uint(ta.Id))
	env["algo_config"] = ta.AlgoConfig //下发参数给sidecar
	env["model_config_str"] = ta.Algo.Algorithm.Detail
	env["image_dir"] = strings.Join(cameraImageDir, ";")
	env["IMAGE_PATH"] = r.bc.Data.Kubernetes.FrameTaskBaseDir
	env["camera_box"] = string(boxJson)
	env["face_group_name"] = extra.FaceGroupName
	env["model_interval"] = fmt.Sprintf("%.3f", interval)
	env["mq_url"] = r.bc.Data.Mq.Url
	env["mq_queue_name"] = ta.Algo.Algorithm.AppName
	env["mq_exchange_name"] = r.bc.Data.Mq.ExchangeName
	env["mq_exchange_type"] = r.bc.Data.Mq.ExchangeType
	env["redis_url"] = r.bc.Data.Redis.Addr + "," + strconv.Itoa(int(r.bc.Data.Redis.Db)) + "," + r.bc.Data.Redis.Password

	mkdirCameraList := make([]string, 0)
	for _, dir := range cameraImageDir {
		mkdirCameraList = append(mkdirCameraList, fmt.Sprintf("mkdir -p %s", dir))
	}

	cameraConf := r.genCameraConf(ta, interval)
	env["CAMERA_CONF"] = string(cameraConf)

	param := FrameJobParam{
		Task:           ta,
		Env:            env,
		FfmpegCommand:  []string{"/usr/local/bin/run_ffmpeg"},
		SideCarCommand: []string{"/bin/sh", "-c", "./sidecar_new"},
		InitCommand:    []string{"/bin/sh", "-c", strings.Join(mkdirCameraList, " && ")},
		CameraCount:    len(ta.Cameras),
	}

	return r.startFrameTaskPod(ctx, param)
}

func ToK8sEnvVars(e map[string]string) []v1.EnvVar {
	var env []v1.EnvVar
	for k, v := range e {
		env = append(env, v1.EnvVar{
			Name:  k,
			Value: v,
		})
	}
	return env
}

// 启动图片帧任务
func (r *scheduler) startFrameTaskPod(ctx context.Context, jobParam FrameJobParam) error {
	// 定义 init 容器
	initContainer := v1.Container{
		Name:    "initdir",
		Image:   r.bc.Data.Kubernetes.GetBusyboxImage(),
		Command: jobParam.InitCommand,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "image-dir",
				MountPath: r.bc.Data.Kubernetes.FrameTaskBaseDir,
			},
		},
	}

	// 定义 ffmpeg 容器
	ffmpegContainer := v1.Container{
		Name:    "ffmpeg",
		Image:   r.bc.Data.Kubernetes.GetFrameFfmpegImage(),
		Command: jobParam.FfmpegCommand,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "image-dir",
				MountPath: r.bc.Data.Kubernetes.FrameTaskBaseDir,
			},
		},
		Env: ToK8sEnvVars(jobParam.Env),
	}
	// 当Interval>=125,即截帧频率小于5秒一帧时，使用定时任务截帧，不使用GPU加速
	if jobParam.Task.Algo.Interval < 125 {
		// 计算需要的资源
		containerResource := k8s.GetTaskResourceLimits(k8s.FrameAlgoCudaFactor, jobParam.Task.AlgoConfig, jobParam.Task.Algo.Interval, jobParam.CameraCount, r.bc.Data.Kubernetes.GpuResourceKey, r.bc.Data.Kubernetes.GpuOpenVirtual)

		ffmpegContainer.Resources = v1.ResourceRequirements{
			Limits:   containerResource,
			Requests: containerResource,
		}
	}

	// 定义 sidecar 容器
	sidecarContainer := v1.Container{
		Name:    "sidecar",
		Image:   r.bc.Data.Kubernetes.GetFrameSidecarImage(),
		Command: jobParam.SideCarCommand,
		VolumeMounts: []v1.VolumeMount{
			{
				Name:      "image-dir",
				MountPath: r.bc.Data.Kubernetes.FrameTaskBaseDir,
			},
		},
		Env: ToK8sEnvVars(jobParam.Env),
		Resources: v1.ResourceRequirements{
			Requests: v1.ResourceList{
				v1.ResourceEphemeralStorage: resource.MustParse("100Mi"),
			},
		},
	}

	// 定义 Pod
	pod := &k8sCoreV1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName(jobParam.Task.Id),
			Namespace: r.bc.Data.Kubernetes.Namespace,
			Labels: map[string]string{
				"app": TaskPodLabel,
			},
		},
		Spec: k8sCoreV1.PodSpec{
			NodeSelector: map[string]string{"node-type": jobParam.Task.Device.Model}, // 使用前端传递的node标签
			InitContainers: []v1.Container{
				initContainer,
			},
			Containers: []v1.Container{
				ffmpegContainer,
				sidecarContainer,
			},
			RestartPolicy: v1.RestartPolicyNever,
			Volumes: []v1.Volume{
				{
					Name: "image-dir",
					VolumeSource: v1.VolumeSource{
						EmptyDir: &v1.EmptyDirVolumeSource{},
					},
				},
			},
		},
	}
	// 设置 HostAliases
	hostAliases := r.bc.Data.Kubernetes.HostAliases
	if len(hostAliases) > 0 {
		r.log.Debugf("add hostAliases: %v", hostAliases)
		setHostAliases(&pod.Spec, hostAliases)
	}
	// 启动 Pod
	_, err := r.client.CoreV1().Pods(r.bc.Data.Kubernetes.Namespace).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		r.log.Errorf("failed to create pod: %s", err)
		return err
	}

	return nil
}
func setHostAliases(podSpec *v1.PodSpec, hostAliases []*conf.Data_HostAlias) {
	// 将 HostAliases 转换为 Kubernetes HostAlias
	k8sHostAliases := make([]v1.HostAlias, 0)
	for _, ha := range hostAliases {
		k8sHostAlias := v1.HostAlias{
			IP:        ha.Ip,
			Hostnames: ha.Hostnames,
		}
		k8sHostAliases = append(k8sHostAliases, k8sHostAlias)
	}
	// 设置 HostAliases 到 PodSpec
	podSpec.HostAliases = k8sHostAliases
}
func configMapName(taskId uint64) string {
	return fmt.Sprintf("config-map-%d", taskId)
}

// 图片帧和视频流类任务统一命名
func podName(taskId uint64) string {
	return fmt.Sprintf("ias-task-%d", taskId)
}

func containerName(ta *biz.Task) string {
	switch ta.Type {
	case biz.TaskTypeStream:
		return fmt.Sprintf("ias-deepstream-container-%d", ta.Id)
	case biz.TaskTypeFrame: // 主容器
		return "sidecar"
	}
	return ""
}

// createConfigMap 创建配置文件
func (r *scheduler) createConfigMap(ctx context.Context, ta *biz.Task, conf *agent.Config) error {
	// 序列化配置文件
	confData, err := yaml.Marshal(conf)
	if err != nil {
		r.log.Errorf("failed to marshal conf: %s", err)
		return err
	}

	// 定义 configmap
	cm := &k8sCoreV1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName(ta.Id),
			Namespace: r.bc.Data.Kubernetes.Namespace,
		},
		Data: map[string]string{
			"config.yaml": string(confData),
		},
	}

	// 创建 configmap
	_, err = r.client.CoreV1().ConfigMaps(r.bc.Data.Kubernetes.Namespace).Create(ctx, cm, metav1.CreateOptions{})
	if err != nil {
		r.log.Errorf("failed to create configmap: %s", err)
		return err
	}

	return nil
}

// deleteConfigMap 删除配置文件
func (r *scheduler) deleteConfigMap(ctx context.Context, ta *biz.Task) error {
	err := r.client.CoreV1().ConfigMaps(r.bc.Data.Kubernetes.Namespace).Delete(ctx, configMapName(ta.Id), metav1.DeleteOptions{})
	if err != nil {
		r.log.Errorf("failed to delete configmap: %s", err)
		return err
	}
	return nil
}

func (r *scheduler) ensureConfigMap(ctx context.Context, ta *biz.Task, conf *agent.Config) error {
	// 检查是否已存在 configmap，已有的话先删除
	cm, err := r.client.CoreV1().ConfigMaps(r.bc.Data.Kubernetes.Namespace).Get(ctx, configMapName(ta.Id), metav1.GetOptions{})
	if err != nil && !k8sErrors.IsNotFound(err) {
		r.log.Errorf("failed to find configmap: %s", err)
		return err
	}

	if cm != nil && err == nil {
		err = r.deleteConfigMap(ctx, ta)
		if err != nil {
			r.log.Errorf("failed to delete configmap: %s", err)
			return err
		}
	}

	return r.createConfigMap(ctx, ta, conf)
}

// startStreamTask 启动视频流任务
func (r *scheduler) startStreamTask(ctx context.Context, ta *biz.Task, conf *agent.Config) error {
	// 创建 configmap
	if err := r.ensureConfigMap(ctx, ta, conf); err != nil {
		r.log.Errorf("failed to create configmap: %s", err)
		return err
	}

	// 启动算法参数
	algoArgs := []string{
		" /app/ias-agent -conf /conf/agent/config.yaml",
	}

	// 计算需要的资源
	containerResource := k8s.GetTaskResourceLimits(k8s.StreamAlgoCudaFactor, ta.AlgoConfig, ta.Algo.Interval, len(ta.Cameras), r.bc.Data.Kubernetes.GpuResourceKey, r.bc.Data.Kubernetes.GpuOpenVirtual)

	// 定义 Pod
	pod := &k8sCoreV1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName(ta.Id),
			Namespace: r.bc.Data.Kubernetes.Namespace,
			Labels: map[string]string{
				"app": TaskPodLabel,
			},
		},
		Spec: k8sCoreV1.PodSpec{
			NodeSelector:  map[string]string{"node-type": ta.Device.Model},
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []k8sCoreV1.Container{
				{
					Name:  containerName(ta),
					Image: ta.Algo.Algorithm.Image,
					VolumeMounts: []k8sCoreV1.VolumeMount{
						{
							Name:      "config-volume",
							MountPath: "/conf/agent",
							ReadOnly:  true,
						},
					},
					Command: []string{"/bin/sh", "-c"},
					Args:    algoArgs,
					Resources: v1.ResourceRequirements{
						Limits:   containerResource,
						Requests: containerResource,
					},
				},
			},
			Volumes: []k8sCoreV1.Volume{
				{
					Name: "config-volume",
					VolumeSource: k8sCoreV1.VolumeSource{
						ConfigMap: &k8sCoreV1.ConfigMapVolumeSource{
							LocalObjectReference: k8sCoreV1.LocalObjectReference{
								Name: configMapName(ta.Id),
							},
						},
					},
				},
			},
		},
	}
	// 设置 HostAliases
	hostAliases := r.bc.Data.Kubernetes.HostAliases
	if len(hostAliases) > 0 {
		r.log.Debugf("add hostAliases: %v", hostAliases)
		setHostAliases(&pod.Spec, hostAliases)
	}
	// 启动 Pod
	_, err := r.client.CoreV1().Pods(r.bc.Data.Kubernetes.Namespace).Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		r.log.Errorf("failed to create pod: %s", err)
		defer r.deleteConfigMap(ctx, ta) // 创建失败时删除 configmap
		return err
	}

	r.log.Info("Pod started successfully!")

	return nil
}

// getTaskResourceLimits 获取任务启动需要申请的资源
func (r *scheduler) getTaskResourceLimits(isStream bool, memory, cameraNum int64, interval, factor float64) v1.ResourceList {
	gpuKey := r.bc.Data.Kubernetes.GpuResourceKey
	list := v1.ResourceList{
		v1.ResourceName(gpuKey + "gpu"): *resource.NewQuantity(1, resource.DecimalSI),
	}
	if !r.bc.Data.Kubernetes.GpuOpenVirtual {
		return list
	}

	if isStream {
		memory = 3000 + 200*cameraNum
	}
	vcore := math.Ceil(factor * interval * float64(cameraNum))
	list[v1.ResourceName(gpuKey+"vcuda-core")] = *resource.NewQuantity(int64(vcore), resource.DecimalSI)
	list[v1.ResourceName(gpuKey+"vcuda-memory")] = *resource.NewQuantity(memory, "")

	return list
}

// Stop 停止任务
func (r *scheduler) Stop(ctx context.Context, ta *biz.Task) error {
	podName := podName(ta.Id)
	err := r.client.CoreV1().Pods(r.bc.Data.Kubernetes.Namespace).Delete(ctx, podName, metav1.DeleteOptions{})
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			r.log.Errorf("Pod %s in namespace %s not found\n", podName, r.bc.Data.Kubernetes.Namespace)
		} else {
			r.log.Errorf("failed to delete pod: %s", err)
		}
		return err
	}

	if ta.Type == biz.TaskTypeStream {
		if err := r.deleteConfigMap(ctx, ta); err != nil {
			r.log.Errorf("failed to delete configmap: %s", err)
			return err
		}
	}

	// 加入redis列表，告诉任务正在停止中，提供 stopping 状态
	r.rdb.SAdd(context.Background(), redisKeyJobCache, fmt.Sprintf("%d", ta.Id))
	r.log.Debugf("task %d add to redis stopping", ta.Id)

	r.log.Info("Pod stopped successfully!")

	return nil
}

// GetLog 查询任务日志
func (r *scheduler) GetLog(ctx context.Context, ta *biz.Task, conn *websocket.Conn) error {
	podName := podName(ta.Id)
	req := r.client.CoreV1().Pods(r.bc.Data.Kubernetes.Namespace).GetLogs(podName, &v1.PodLogOptions{
		Container: containerName(ta),
		Follow:    true,
	})
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			r.log.Errorf("Pod %s in namespace %s not found\n", podName, r.bc.Data.Kubernetes.Namespace)
		}
		if statusError, isStatus := err.(*k8sErrors.StatusError); isStatus {
			r.log.Errorf("Error getting logs for Pod %s in namespace %s: %v\n", podName, r.bc.Data.Kubernetes.Namespace, statusError.ErrStatus.Message)
		}
		return err
	}

	for {
		buf := make([]byte, 1024)

		n, err := podLogs.Read(buf)
		if err != nil {
			if err != io.EOF {
				r.log.Errorf("failed to read log: %s", err)
				return err
			}
			r.log.Errorf("read EOF: %s", err)
			break
		}

		var encodedMsg = make([]byte, base64.StdEncoding.EncodedLen(len(buf[:n])))
		base64.StdEncoding.Encode(encodedMsg, buf[:n])

		if err := conn.WriteMessage(websocket.TextMessage, encodedMsg); err != nil {
			r.log.Errorf("write data to ws failed %s", err)
			return err
		}
	}

	return nil
}

func getTaskStatusFromPod(pod v1.Pod) (biz.TaskStatus, string) {
	switch pod.Status.Phase {
	case v1.PodPending:
		return biz.TaskStatusStarting, "任务启动中"
	case v1.PodRunning:
		for _, c := range pod.Status.ContainerStatuses {
			// 存在一个pod多个container中，某个container失败， 但是总体还是running状态
			if st := c.State.Terminated; st != nil {
				return biz.TaskStatusFailed, "任务运行失败"
			}
		}
		return biz.TaskStatusRunning, "任务运行中"
	case v1.PodSucceeded: // 正常停止
		return biz.TaskStatusStopped, "任务已停止"
	case v1.PodFailed:
		return biz.TaskStatusFailed, "任务运行失败"
	case v1.PodUnknown:
		return biz.TaskStatusUnknown, "任务状态未知"
	}
	return biz.TaskStatusUnknown, "任务状态未知"
}

func populateTasksWithStatus(taskIdList []uint, status biz.TaskStatus) []biz.TaskIdStatus {
	var statuses = make([]biz.TaskIdStatus, len(taskIdList))
	for i := range taskIdList {
		statuses[i] = biz.TaskIdStatus{
			TaskId: taskIdList[i],
			Status: status,
			Reason: "获取任务状态失败，任务状态未知",
		}
	}
	return statuses
}

func (r *scheduler) GetStatuses(ctx context.Context, taskIdList []uint) ([]biz.TaskIdStatus, error) {
	// 获取集群中所有任务 Pods
	list, err := r.client.CoreV1().Pods(r.bc.Data.Kubernetes.Namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: TaskPodLabelSelector,
	})
	if err != nil {
		r.log.Errorf("failed to get pods: %s", err)
		return populateTasksWithStatus(taskIdList, biz.TaskStatusUnknown), err
	}

	var pods = make(map[string]v1.Pod)
	for _, pod := range list.Items {
		pods[pod.Name] = pod
	}

	// 计算任务状态
	var statuses []biz.TaskIdStatus
	var recheckingTasks = make(map[string]int) // 需要再次检查的状态，task id -> array index
	var allTasks = make(map[string]biz.TaskStatus)
	for i, taskId := range taskIdList {
		status := biz.TaskIdStatus{
			TaskId: taskId,
		}
		if pod, ok := pods[podName(uint64(taskId))]; ok {
			status.Status, status.Reason = getTaskStatusFromPod(pod)
			if status.Status == biz.TaskStatusRunning || status.Status == biz.TaskStatusFailed {
				recheckingTasks[fmt.Sprintf("%d", status.TaskId)] = i
			}
		} else { // 查不到时认为已停止
			status.Status = biz.TaskStatusStopped
			status.Reason = "任务已停止"
		}

		allTasks[fmt.Sprintf("%d", status.TaskId)] = status.Status

		statuses = append(statuses, status)
	}

	// 处理停止中任务
	stoppingTasks, err := r.rdb.SMembers(context.Background(), redisKeyJobCache).Result()
	if err != nil {
		r.log.Errorf("failed to get stopping tasks: %s", err)
		return statuses, err
	}
	for _, task := range stoppingTasks {
		// 1. Pod 状态还是 running，实际在停止中  2. 停止中的偶发 failed，返回停止中
		if idx, ok := recheckingTasks[task]; ok {
			statuses[idx].Status = biz.TaskStatusStopping
			statuses[idx].Reason = "任务停止中"
			r.log.Debugf("[GetStatuses] task %s is stopping", task)
		} else if status, ok2 := allTasks[task]; ok2 && status == biz.TaskStatusStopped {
			r.log.Debugf("[GetStatuses] task %s delete stopping from redis", task)
			r.rdb.SRem(context.Background(), redisKeyJobCache, task)
		}
	}

	return statuses, nil
}

func getVcoreFactor(defaultVal float64, algoExtend string) float64 {
	factor := defaultVal
	if algoExtend != "" {
		var extend map[string]interface{}
		err := json.Unmarshal([]byte(algoExtend), &extend)
		if err == nil {
			if val, ok := extend["factor"].(float64); ok && val > 0 {
				factor = val
			}
		}
	}
	return factor
}
