package k8s

import (
	"encoding/json"
	"math"

	"github.com/blues120/ias-core/pkg/convert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	CpuResourceKey    = "cpu"
	MemoryResourceKey = "memory"
)

type CudaFactor struct {
	CudaMemoryBase   float64 `json:"cudaMemoryBase"`
	CudaMemoryFactor float64 `json:"cudaMemoryFactor"`
	CudaCoreBase     float64 `json:"cudaCoreBase"`
	CudaCoreFactor   float64 `json:"cudaCoreFactor"`
}

var (
	// 图片帧算法倍率因子
	FrameAlgoCudaFactor = CudaFactor{
		CudaMemoryBase:   0,
		CudaMemoryFactor: 350,
		CudaCoreBase:     0,
		CudaCoreFactor:   0.4,
	}
	// 视频流算法倍率因子
	StreamAlgoCudaFactor = CudaFactor{
		CudaMemoryBase:   3000,
		CudaMemoryFactor: 200,
		CudaCoreBase:     0,
		CudaCoreFactor:   0.8,
	}
)

func GetTaskResourceLimits(cudaFactor CudaFactor, algoConfig string, algoInterval float64, cameraNum int, gpuKey string, openVirtualGpu bool) v1.ResourceList {
	resourceKeys := GetResourceKeys(gpuKey)

	list := v1.ResourceList{
		v1.ResourceName(resourceKeys.Gpu): *resource.NewQuantity(1, resource.DecimalSI),
	}
	if !openVirtualGpu {
		return list
	}

	var extend map[string]interface{}
	err := json.Unmarshal([]byte(algoConfig), &extend)
	if err == nil {
		if val, ok := extend["cudaMemoryBase"].(float64); ok {
			cudaFactor.CudaMemoryBase = val
		}
		if val, ok := extend["cudaMemoryFactor"].(float64); ok {
			cudaFactor.CudaMemoryFactor = val
		}
		if val, ok := extend["cudaCoreBase"].(float64); ok {
			cudaFactor.CudaCoreBase = val
		}
		if val, ok := extend["cudaCoreFactor"].(float64); ok {
			cudaFactor.CudaCoreFactor = val
		}
	}

	fps := convert.ConvertDropFrameIntervalToFPS(int(algoInterval))
	vcudaMemory := cudaFactor.CudaMemoryBase + cudaFactor.CudaMemoryFactor*float64(cameraNum)          // base + factor * cameraNum
	vcudaCore := math.Ceil(cudaFactor.CudaCoreBase + cudaFactor.CudaCoreFactor*float64(cameraNum)*fps) // base + factor * cameraNum * fps

	list[v1.ResourceName(resourceKeys.GpuVCore)] = *resource.NewQuantity(int64(vcudaCore), resource.DecimalSI)
	list[v1.ResourceName(resourceKeys.GpuVMemory)] = *resource.NewQuantity(int64(vcudaMemory), "")

	return list
}

type ResourceKeys struct {
	Cpu        string
	Memory     string
	Gpu        string
	GpuVCore   string
	GpuVMemory string
}

func GetResourceKeys(gpuResourceKey string) *ResourceKeys {
	return &ResourceKeys{
		Cpu:        CpuResourceKey,
		Memory:     MemoryResourceKey,
		Gpu:        gpuResourceKey + "gpu",
		GpuVCore:   gpuResourceKey + "vcuda-core",
		GpuVMemory: gpuResourceKey + "vcuda-memory",
	}
}
