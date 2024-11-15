package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/tasklimits"
)

type taskLimitsRepo struct {
	data *Data

	log *log.Helper
}

func NewTaskLimitsRepo(data *Data, logger log.Logger) biz.TasklimitsRepo {
	return &taskLimitsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// TaskLimitsEntToBiz 将ent中的tasklimits模型转换为biz中的TaskLimits模型
func TaskLimitsEntToBiz(a *ent.TaskLimits) *biz.TaskLimits {
	tasklimits := &biz.TaskLimits{
		ID:            a.ID,
		Model:         a.Model,
		MaxCameraNum:  a.MaxCameraNum,
		AlgoNum:       a.AlgoNum,
		MaxSubTaskNum: a.MaxSubTaskNum,
	}
	return tasklimits
}

// TaskLimitsBizToEnt 将biz中的TaskLimits模型转换为ent中的TaskLimits模型
func TaskLimitsBizToEnt(a *biz.TaskLimits) *ent.TaskLimits {
	tasklimits := &ent.TaskLimits{
		ID:            a.ID,
		Model:         a.Model,
		MaxCameraNum:  a.MaxCameraNum,
		AlgoNum:       a.AlgoNum,
		MaxSubTaskNum: a.MaxSubTaskNum,
	}
	return tasklimits
}

func (r *taskLimitsRepo) List(ctx context.Context, filter *biz.TaskListFilter) ([]*biz.TaskLimits, int, error) {
	query := r.data.db.TaskLimits(ctx).Query()

	// 查询总条数
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	// 组装分页条件
	if filter.Pagination != nil {
		query = query.Offset(filter.Pagination.Offset()).Limit(filter.Pagination.PageSize)
	}
	// 执行查询
	results, err := query.All(ctx)
	if err != nil {
		return nil, 0, err
	}
	// 转换为业务层对象
	list := make([]*biz.TaskLimits, len(results))
	for i, result := range results {
		list[i] = TaskLimitsEntToBiz(result)
	}
	return list, total, nil
}

func (r *taskLimitsRepo) Save(ctx context.Context, taskLimits *biz.TaskLimits) (uint64, error) {
	data, err := r.data.db.TaskLimits(ctx).Create().SetTaskLimits(TaskLimitsBizToEnt(taskLimits)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *taskLimitsRepo) FindByModel(ctx context.Context, model string) (*biz.TaskLimits, error) {
	foundInform, err := r.data.db.TaskLimits(ctx).Query().Where(tasklimits.Model(model)).First(ctx)
	if err != nil {
		return nil, err
	}
	return TaskLimitsEntToBiz(foundInform), nil
}

func (r *taskLimitsRepo) Update(ctx context.Context, model string, taskLimits *biz.TaskLimits) error {
	return r.data.db.TaskLimits(ctx).Update().Where(tasklimits.ModelEQ(model)).SetTaskLimits(TaskLimitsBizToEnt(taskLimits)).Exec(ctx)
}

func (r *taskLimitsRepo) DeleteByModel(ctx context.Context, model string) (int, error) {
	return r.data.db.TaskLimits(ctx).Delete().Where(tasklimits.ModelEQ(model)).Exec(ctx)
}
