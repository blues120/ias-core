//go:build skip
// +build skip

package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/pkg/log/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func initMetricsUc() *K8sMetricsUsecase {
	logger := zap.NewLogger(&conf.Log{Mode: "console", Level: "debug"})
	uc := NewK8sMetricsUsecase(&conf.Bootstrap{
		Data: &conf.Data{
			Kubernetes: &conf.Data_Kubernetes{
				KubeConfig: "./.kube/testconfig",
			},
		},
	}, logger)
	return uc
}

func TestNodeAllocatableResources(t *testing.T) {
	uc := initMetricsUc()
	resources, err := uc.NodeAllocatableResources(context.Background(), v1.ListOptions{
		LabelSelector: "node-type=a10",
	})
	assert.Equal(t, nil, err)
	resourcesBytes, _ := json.Marshal(resources)
	fmt.Println(string(resourcesBytes))
}

func TestNodeRequestedResources(t *testing.T) {
	uc := initMetricsUc()
	resources, err := uc.NodeRequestedResources(context.Background(), v1.ListOptions{
		LabelSelector: "node-type=a10",
	})
	assert.Equal(t, nil, err)
	resourcesBytes, _ := json.Marshal(resources)
	fmt.Println(string(resourcesBytes))
}

func TestNodeAvailableResources(t *testing.T) {
	uc := initMetricsUc()
	resources, err := uc.NodeAvailableResources(context.Background(), v1.ListOptions{
		LabelSelector: "node-type=a10",
	})
	assert.Equal(t, nil, err)
	resourcesBytes, _ := json.Marshal(resources)
	fmt.Println(string(resourcesBytes))
}
