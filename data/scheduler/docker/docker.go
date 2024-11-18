package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/data/scheduler/agent"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/file"
	"gopkg.in/yaml.v3"
)

type scheduler struct {
	bc *conf.Bootstrap

	log *log.Helper
}

func NewScheduler(bc *conf.Bootstrap, logger log.Logger) biz.SchedulerRepo {
	if bc == nil || logger == nil {
		return nil
	}

	return &scheduler{
		bc:  bc,
		log: log.NewHelper(logger),
	}
}

// Close 关闭
func (r *scheduler) Close() error {
	return nil
}

// Start 启动任务
func (r *scheduler) Start(ctx context.Context, ta *biz.Task) error {
	config := agent.TaskToAgentConfig(r.bc, ta, r.bc.AgentExtend)
	payload, _ := json.Marshal(config)
	r.log.Debugf("docker Start:%+v, %s\n", config, string(payload))
	switch ta.Type {
	case biz.TaskTypeFrame:
		return fmt.Errorf("暂未实现本地运行图片帧任务")
	case biz.TaskTypeStream:
		return r.startStreamTask(ctx, ta, config)
	default:
		return fmt.Errorf("unsupported task type: %v", ta.Type)
	}
}

// Stop 停止任务
// 1. 需要清理本地目录的所有资源。（如果重新启动必须重新拉起）
func (r *scheduler) Stop(ctx context.Context, ta *biz.Task) error {
	if ta.Status == biz.TaskStatusRunning {
		containerId := agent.GetTaskId(ta)
		return r.stopContainer(fmt.Sprintf("ias_deepstream_%s", containerId), false)
	}
	return errors.ErrorTaskStopError("任务未运行")
}

// GetLog 查询任务日志
func (r *scheduler) GetLog(ctx context.Context, ta *biz.Task, conn *websocket.Conn) error {
	// TODO implement
	panic("implement me")
}

// GetStatus 查询任务状态
func (r *scheduler) GetStatuses(ctx context.Context, taskIdList []uint) ([]biz.TaskIdStatus, error) {
	// TODO implement
	panic("implement me")
}

func (r *scheduler) startStreamTask(ctx context.Context, ta *biz.Task, conf *agent.Config) error {
	containerId := agent.GetTaskId(ta)
	configFile := fmt.Sprintf("config-%s.yaml", containerId)
	// 启动算法
	algoArgs := []string{
		fmt.Sprintf(" cd /root/ibox/deepstream-app && /app/ias-agent -conf /conf/%s", configFile),
	}

	localConfigPath := "/root/.ibox/config"
	if _, err := file.EnsureDir(localConfigPath); err != nil {
		return err
	}

	if err := r.saveObjectToPath(conf, localConfigPath, configFile); err != nil {
		r.log.Error(fmt.Sprintf("Generate Agent Config error %v", err))
		return err
	}

	algoImageName := ta.Algo.Algorithm.Image
	algoContainerName := fmt.Sprintf("ias_deepstream_%s", containerId)

	mounts := []mount.Mount{
		{
			Type:   "bind",
			Source: conf.Agent.Data.Oss.Local.StorePath,
			Target: conf.Agent.Data.Oss.Local.StorePath,
		},
		{
			Type:   "bind",
			Source: fmt.Sprintf("%s/%s", localConfigPath, configFile),
			Target: fmt.Sprintf("/conf/%s", configFile),
		},
		{
			Type:   "bind",
			Source: "/data/image",
			Target: "/data/image",
		},
	}

	err := r.runContainer(algoImageName, algoContainerName, algoArgs, mounts)
	if err != nil {
		// 删除旧的
		_ = r.stopContainer(algoContainerName, true)
	}
	return err
}

func (r *scheduler) runContainer(imageName string, containerName string, args []string, mounts []mount.Mount) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()

	if err := r.checkLicenseImageExistsIfNeed(ctx, cli, imageName); err != nil {
		return err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   append([]string{"bash", "-c"}, args...),
	}, &container.HostConfig{
		// AutoRemove: true, // 不能和restart policy 一同使用
		RestartPolicy: container.RestartPolicy{
			Name: "unless-stopped",
		},
		NetworkMode: "host",
		Privileged:  false,
		Runtime:     "nvidia",
		Mounts:      mounts,
	}, nil, nil, containerName)
	r.log.Info("create container--->", resp, err, imageName, args)
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		r.log.Error("start container failed ======> ", err)
		return err
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		r.log.Error("--------> get cli container log failed", err)
		return err
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return nil
}

func (r *scheduler) stopContainer(containerName string, all bool) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Limit:   1,
		All:     all,
		Filters: filters.NewArgs(filters.Arg("name", containerName)),
	})

	if err != nil || len(containers) == 0 {
		return err
	}

	containerId := containers[0].ID
	return cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{
		// RemoveVolumes: true,
		Force: true,
	})
}

func (r *scheduler) saveObjectToPath(config *agent.Config, path string, name string) (err error) {
	data, err := yaml.Marshal(&config)
	if err != nil {
		r.log.Errorf("yaml marshal err: %v", err)
		return err
	}

	if err = os.WriteFile(path+"/"+name, data, 0); err != nil {
		r.log.Errorf("Write file error %v", err)
	}
	return nil
}

// checkLicenseImageExistsIfNeed 检查是否为license镜像
func (r *scheduler) checkLicenseImageExistsIfNeed(ctx context.Context, cli *client.Client, imageName string) error {
	// license镜像需要rdac-server作为辅助
	if !strings.Contains(imageName, "-lic") {
		return nil
	}

	// 先确认有没有rdac-server这个容器
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Limit:   1,
		Filters: filters.NewArgs(filters.Arg("name", "rdac-server")),
	})
	if err != nil {
		return err
	}
	if len(containers) == 0 {
		// 如果rdac-server没有启动，请先启动
		// 启动指令:docker run -d --restart=always --net=host -v /opt/ctyun:/opt/ctyun --name=rdac-server harbor.ctyuncdn.cn/ai-service/rdac/rdac-server:aarch64-0.1.0
		r.log.Info("no rdac-server container")
		return errors.ErrorInvalidParam("请先启动rdac-server")
	}
	return nil
}
