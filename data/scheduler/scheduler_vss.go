package scheduler

import (
	"fmt"

	"github.com/blues120/ias-core/data/scheduler/ecx"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
)

func NewSchedulerVSSRepo(bc *conf.Bootstrap, signUc *biz.SignatureUsecase, logger log.Logger) (biz.SchedulerVSSRepo, error) {
	switch bc.Scheduler.Mode {
	case conf.Scheduler_ecx_vss:
		return ecx.NewSchedulerVSS(bc, signUc, logger), nil
	case conf.Scheduler_docker, conf.Scheduler_ecx, conf.Scheduler_k8s, conf.Scheduler_sophgo, conf.Scheduler_mix:
		return nil, nil
	default:
		return nil, fmt.Errorf("unsupported scheduler mode: %s", bc.Scheduler.Mode)
	}
}
