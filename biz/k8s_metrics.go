package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	mc "k8s.io/metrics/pkg/client/clientset/versioned"

	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/k8s"
)

type K8sMetricsUsecase struct {
	metricsCli *mc.Clientset         // 用于获取系统实时使用率信息，需要在 k8s 集群中部署 metrics-server，暂不使用
	coreCli    *kubernetes.Clientset // 用于获取系统静态使用率信息，比如已申请资源比例

	cfg *conf.Data_Kubernetes

	log *log.Helper
}

func NewK8sMetricsUsecase(bc *conf.Bootstrap, logger log.Logger) *K8sMetricsUsecase {
	// 指定配置文件
	config, err := clientcmd.BuildConfigFromFlags("", bc.Data.Kubernetes.KubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建 metrics client
	mcs, err := mc.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 创建 core client
	ccs, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &K8sMetricsUsecase{
		metricsCli: mcs,
		coreCli:    ccs,
		cfg:        bc.Data.Kubernetes,
		log:        log.NewHelper(logger),
	}
}

type NodeAllocatableResource struct {
	// 节点名
	NodeName string
	/*
		集群总资源容量，返回数值以方便后续计算
		对 cpu，返回单位为核数；对 memory，返回单位为 Byte; 对 gpu vcuda-memory，返回单位为 MB
	*/
	Capacity map[string]*resource.Quantity
	/*
		用户可用资源量，去除了 k8s 本身占用的部分资源等，返回数值以方便后续计算。通常应使用 Allocatable 计算资源使用率
		对 cpu，返回单位为核数；对 memory，返回单位为 Byte；对 gpu vcuda-memory，返回单位为 MB
	*/
	Allocatable map[string]*resource.Quantity
}

/*
系统可分配给用户进程的资源
示例输出:

	{
	    "Capacity": {
	        "cpu": 32,
	        "ephemeral-storage": 83053432832,
	        "hugepages-1Gi": 0,
	        "hugepages-2Mi": 0,
	        "ideal.com/gpu": 20,
	        "ideal.com/vcuda-core": 400,
	        "ideal.com/vcuda-memory": 90927,
	        "memory": 67547693056,
	        "nvidia.com/gpu": 4,
	        "pods": 110
	    },
	    "Allocatable": {
	        "cpu": 32,
	        "ephemeral-storage": 74748089426,
	        "hugepages-1Gi": 0,
	        "hugepages-2Mi": 0,
	        "ideal.com/gpu": 20,
	        "ideal.com/vcuda-core": 400,
	        "ideal.com/vcuda-memory": 90927,
	        "memory": 67442835456,
	        "nvidia.com/gpu": 4,
	        "pods": 110
	    }
	}

options 用来过滤节点
*/
func (uc *K8sMetricsUsecase) NodeAllocatableResources(ctx context.Context, options v1.ListOptions) ([]*NodeAllocatableResource, error) {
	nodes, err := uc.coreCli.CoreV1().Nodes().List(ctx, options)
	if err != nil {
		uc.log.Errorf("failed to list all nodes: %s", err)
		return nil, err
	}

	var ret = make([]*NodeAllocatableResource, 0)
	for i := range nodes.Items {
		var resources = NodeAllocatableResource{
			NodeName:    nodes.Items[i].Name,
			Capacity:    make(map[string]*resource.Quantity),
			Allocatable: make(map[string]*resource.Quantity),
		}
		for k, v := range nodes.Items[i].Status.Capacity {
			val := v.DeepCopy()
			resources.Capacity[string(k)] = &val
		}
		for k, v := range nodes.Items[i].Status.Allocatable {
			val := v.DeepCopy()
			resources.Allocatable[string(k)] = &val
		}
		ret = append(ret, &resources)
	}
	return ret, nil
}

type NodeRequestedResource struct {
	// 节点名
	NodeName string
	/*
		容器已申请资源
		对 cpu，返回单位为核数；对 memory，返回单位为 Byte; 对 gpu vcuda-memory，返回单位为 MB
	*/
	Requests map[string]*resource.Quantity
	/*
		容器申请资源上限
		对 cpu，返回单位为核数；对 memory，返回单位为 Byte；对 gpu vcuda-memory，返回单位为 MB
	*/
	Limits map[string]*resource.Quantity
}

// 将 new 的同名资源合并到 total 上
func mergeRequestedResources(total map[string]*resource.Quantity, new coreV1.ResourceList) {
	for k, v := range new {
		if _, ok := total[string(k)]; !ok {
			total[string(k)] = &resource.Quantity{}
		}
		total[string(k)].Add(v)
	}
}

// 将 pod 的申请资源合并到当前总申请资源上
func mergePodRequestedResources(currentTotalRequested map[string]NodeRequestedResource, pod coreV1.Pod) {
	nodeName := pod.Spec.NodeName
	if _, ok := currentTotalRequested[nodeName]; !ok {
		currentTotalRequested[nodeName] = NodeRequestedResource{
			NodeName: nodeName,
			Requests: make(map[string]*resource.Quantity),
			Limits:   make(map[string]*resource.Quantity),
		}
	}

	for i := range pod.Spec.Containers {
		mergeRequestedResources(currentTotalRequested[nodeName].Requests, pod.Spec.Containers[i].Resources.Requests)
		mergeRequestedResources(currentTotalRequested[nodeName].Limits, pod.Spec.Containers[i].Resources.Limits)
	}
}

func (uc *K8sMetricsUsecase) NodeRequestedResources(ctx context.Context, options v1.ListOptions) ([]*NodeRequestedResource, error) {
	// 获取所有命名空间
	ns, err := uc.coreCli.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		uc.log.Errorf("failed to list namespaces: %s", err)
		return nil, err
	}

	var nodeResourcesMap = make(map[string]NodeRequestedResource)
	// 遍历空间获取所有 pod
	for i := range ns.Items {
		name := ns.Items[i].Name
		pods, err := uc.coreCli.CoreV1().Pods(name).List(ctx, v1.ListOptions{})
		if err != nil {
			uc.log.Errorf("failed to list pods for namespace %s: %s", name, err)
			continue
		}
		for j := range pods.Items {
			// 只统计 Non-terminated 的 pod
			if pods.Items[j].Status.Phase == coreV1.PodFailed || pods.Items[j].Status.Phase == coreV1.PodSucceeded {
				fmt.Println("pod stopped", pods.Items[j].Name)
				continue
			}
			mergePodRequestedResources(nodeResourcesMap, pods.Items[j])
		}
	}

	// 获取所有 nodes，根据 node selector 过滤数据
	nodes, err := uc.coreCli.CoreV1().Nodes().List(ctx, options)
	if err != nil {
		uc.log.Errorf("failed to list all nodes: %s", err)
		return nil, err
	}
	var nodesMap = make(map[string]struct{})
	for i := range nodes.Items {
		nodesMap[nodes.Items[i].Name] = struct{}{}
	}

	var ret = make([]*NodeRequestedResource, 0)
	for k, v := range nodeResourcesMap {
		if _, ok := nodesMap[k]; ok {
			ret = append(ret, &v)
		}
	}

	return ret, nil
}

type AvailableResourcePercent struct {
	Quantity     *resource.Quantity // 剩余资源量
	Percent      float64            // 剩余资源占比
	PercentHuman string             // 剩余资源占比字符串，用于前端显示
}

type NodeAvailableResource struct {
	// 节点名
	NodeName string
	/*
		容器可用资源余量
	*/
	Available map[string]*AvailableResourcePercent
}

func remainingResources(allocatable, requested map[string]*resource.Quantity) map[string]*AvailableResourcePercent {
	var remainResources = make(map[string]*AvailableResourcePercent)

	if requested == nil {
		for k, v := range allocatable {
			remainResources[k] = &AvailableResourcePercent{
				Quantity:     v,
				Percent:      1,
				PercentHuman: "100%",
			}
		}
	} else {
		for k, v := range allocatable {
			if requests, ok := requested[k]; ok {
				remain := v.DeepCopy()
				(&remain).Sub(*requests)
				percent := float64(remain.Value()) / float64(v.Value())
				remainResources[k] = &AvailableResourcePercent{
					Quantity:     &remain,
					Percent:      percent,
					PercentHuman: fmt.Sprintf("%.2f%%", percent*100),
				}
			} else {
				remainResources[k] = &AvailableResourcePercent{
					Quantity:     v,
					Percent:      1,
					PercentHuman: "100%",
				}
			}
		}
	}

	return remainResources
}

// 系统还可以给用户进程分配的资源（用户可用资源余量），区别于 NodeAllocatableResources（用户可用资源总量）
func (uc *K8sMetricsUsecase) NodeAvailableResources(ctx context.Context, options v1.ListOptions) ([]*NodeAvailableResource, error) {
	// 获取可用资源总量
	allocatable, err := uc.NodeAllocatableResources(ctx, options)
	if err != nil {
		uc.log.Errorf("failed to get allocatable resources: %s", err)
		return nil, err
	}

	// 获取用户已申请资源
	requested, err := uc.NodeRequestedResources(ctx, options)
	if err != nil {
		uc.log.Errorf("failed to get requested resources: %s", err)
		return nil, err
	}
	var nodeRequestedMap = make(map[string]*NodeRequestedResource)
	for i := range requested {
		nodeRequestedMap[requested[i].NodeName] = requested[i]
	}

	// 计算资源余量
	var ret = make([]*NodeAvailableResource, 0)
	for i := range allocatable {
		var available = NodeAvailableResource{
			NodeName: allocatable[i].NodeName,
		}

		if _, ok := nodeRequestedMap[allocatable[i].NodeName]; !ok {
			available.Available = remainingResources(allocatable[i].Allocatable, nil)
		} else {
			available.Available = remainingResources(allocatable[i].Allocatable, nodeRequestedMap[allocatable[i].NodeName].Limits) // 使用 limits 计算余量，保守计算
		}

		ret = append(ret, &available)
	}

	return ret, nil
}

// 获取各种资源的 key
func (uc *K8sMetricsUsecase) ResourceKeys() *k8s.ResourceKeys {
	return k8s.GetResourceKeys(uc.cfg.GpuResourceKey)
}
