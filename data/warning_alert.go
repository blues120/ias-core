package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
)

type warningAlertRepo struct {
	data *Data

	log *log.Helper
}

func NewWarningAlertRepo(data *Data, logger log.Logger) biz.WarningAlertRepo {
	return &warningAlertRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *warningAlertRepo) Publish(ctx context.Context, channel string, message any) error {
	return r.data.rdb.Publish(ctx, channel, message).Err()
}

func (r *warningAlertRepo) Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	ch := r.data.rdb.Subscribe(ctx, channel).Channel()
	return ch
}

func (r *warningAlertRepo) XAdd(ctx context.Context, args *redis.XAddArgs) error {
	return r.data.rdb.XAdd(ctx, args).Err()
}

func (r *warningAlertRepo) XRead(ctx context.Context, args *redis.XReadArgs) ([]redis.XStream, error) {
	return r.data.rdb.XRead(ctx, args).Result()
}
