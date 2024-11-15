package scheduler

import (
	"fmt"

	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler/ecx"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler/sophgo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler/docker"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler/kubernetes"
)

func NewSchedulerRepo(bc *conf.Bootstrap, rdb *redis.Client, logger log.Logger) (biz.SchedulerRepo, error) {
	switch bc.Scheduler.Mode {
	case conf.Scheduler_docker:
		return docker.NewScheduler(bc, logger), nil
	case conf.Scheduler_ecx:
		return ecx.NewScheduler(bc, logger), nil
	case conf.Scheduler_k8s:
		return kubernetes.NewScheduler(bc, rdb, logger), nil
	case conf.Scheduler_ecx_vss:
		return nil, nil
	case conf.Scheduler_sophgo:
		return sophgo.NewScheduler(bc, logger), nil
	default:
		return nil, fmt.Errorf("unsupported scheduler mode: %s", bc.Scheduler.Mode)
	}
}
