package scheduler

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/data/scheduler/docker"
	"github.com/blues120/ias-core/data/scheduler/ecx"
	"github.com/blues120/ias-core/data/scheduler/kubernetes"
	"github.com/blues120/ias-core/data/scheduler/sophgo"
)

func NewSchedulerRepoSelector(bc *conf.Bootstrap, rdb *redis.Client, logger log.Logger) (biz.SchedulerRepoSelector, error) {
	// 初始化各种调度器, 配置不正确返回nil
	repoMap := map[string]biz.SchedulerRepo{
		"docker": docker.NewScheduler(bc, logger),
		"api":    sophgo.NewScheduler(bc, logger),
		"ecx":    ecx.NewScheduler(bc, logger),
		"k8s":    kubernetes.NewScheduler(bc, rdb, logger),
	}

	return func(mode string) biz.SchedulerRepo {
		repo, ok := repoMap[mode]
		if !ok {
			panic("invalid scheduler mode: " + mode)
		}
		if repo == nil {
			panic(fmt.Sprintf("scheduler not initialized, please check the config, mode: %s, bc: %+v, logger: %+v", mode, bc, logger))
		}
		return repo
	}, nil
}
