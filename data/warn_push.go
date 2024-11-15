package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/task"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/warnpush"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/warnpushlog"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/convert"
)

type warnPushRepo struct {
	data *Data
	log  *log.Helper
}

// NewWarnPushRepo warnPushRepo
func NewWarnPushRepo(data *Data, logger log.Logger) biz.WarnPushRepo {
	return &warnPushRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// WarnPushEntToBiz 将warn push从ent层转化为biz层结构
func WarnPushEntToBiz(wp *ent.WarnPush) *biz.WarnPush {
	bizPush := &biz.WarnPush{
		Id:            wp.ID,
		Name:          wp.Name,
		Type:          wp.Type,
		Url:           wp.URL,
		Remark:        wp.Remark,
		Mode:          wp.Mode,
		Status:        wp.Status,
		CreatedAt:     wp.CreatedAt,
		UpdatedAt:     wp.UpdatedAt,
		TenantID:      wp.TenantID,
		AccessOrgList: wp.AccessOrgList,
	}

	return bizPush
}

// WarnPushEntArrToBiz 将task数组从ent层转化为biz层结构
func WarnPushEntArrToBiz(arr []*ent.WarnPush) []*biz.WarnPush {
	return convert.ToArr[ent.WarnPush, biz.WarnPush](WarnPushEntToBiz, arr)
}

// WarnPushBizToEnt 将warn_push从biz层转化为ent层结构
func WarnPushBizToEnt(wp *biz.WarnPush) *ent.WarnPush {
	return &ent.WarnPush{
		ID:            wp.Id,
		Name:          wp.Name,
		Type:          wp.Type,
		URL:           wp.Url,
		Remark:        wp.Remark,
		Mode:          wp.Mode,
		Status:        wp.Status,
		CreatedAt:     wp.CreatedAt,
		UpdatedAt:     wp.UpdatedAt,
		TenantID:      wp.TenantID,
		AccessOrgList: wp.AccessOrgList,
	}
}

// Save 创建 task
func (r *warnPushRepo) Save(ctx context.Context, wp *biz.WarnPush) (uint64, error) {
	var taskId uint64

	push := WarnPushBizToEnt(wp)
	dbTask, err := r.data.db.WarnPush(ctx).Create().SetWarnPush(push).Save(ctx)
	if err != nil {
		return taskId, err
	}

	taskId = dbTask.ID

	return taskId, err
}

// Update 更新
func (r *warnPushRepo) Update(ctx context.Context, pushId uint64, ta *biz.WarnPush) (bool, error) {
	_, err := r.get(ctx, pushId)
	if err != nil {
		return false, err
	}

	err = r.data.db.WarnPush(ctx).UpdateOneID(pushId).
		SetWarnPush(WarnPushBizToEnt(ta)).Exec(ctx)

	return true, err
}

// UpdateStatus 更新状态
func (r *warnPushRepo) UpdateStatus(ctx context.Context, pushId uint64, status biz.WarnPushStatus) error {
	return r.data.db.WarnPush(ctx).
		Update().
		Where(warnpush.IDEQ(pushId)).
		SetStatus(status).Exec(ctx)
}

// Delete 删除
func (r *warnPushRepo) Delete(ctx context.Context, pushId uint64) error {
	// 删除任务
	return r.data.db.WarnPush(ctx).DeleteOneID(pushId).Exec(ctx)
}

// Find 查询
func (r *warnPushRepo) Find(ctx context.Context, pushId uint64) (*biz.WarnPush, error) {
	push, err := r.get(ctx, pushId)
	if err != nil {
		return nil, err
	}

	return WarnPushEntToBiz(push), nil
}

// CheckExists 检查 WarnPush 的字段值是否重复
func (r *warnPushRepo) CheckExists(ctx context.Context, req *biz.CheckWarnPushExistsRequest) (bool, error) {
	query := r.query(ctx)
	if req.Name != "" {
		query = query.Where(warnpush.NameEQ(req.Name))
	}
	if req.Url != "" {
		query = query.Where(warnpush.URLEQ(req.Url))
	}
	if req.ExcludeId > 0 {
		query = query.Where(warnpush.IDNEQ(req.ExcludeId))
	}
	return query.Exist(ctx)
}

// get 根据ID获取WarnPush
func (r *warnPushRepo) get(ctx context.Context, pushId uint64) (*ent.WarnPush, error) {
	query := r.data.db.WarnPush(ctx).Query().Where(warnpush.IDEQ(pushId))

	return query.First(ctx)
}

// List 批量查询
func (r *warnPushRepo) List(ctx context.Context, filter *biz.PushListFilter) ([]*biz.WarnPush, int, error) {
	query := r.query(ctx)
	if filter != nil {
		query = r.buildQueryByFilter(query, filter)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 分页
	if filter.Pagination != nil {
		query = query.Clone().Offset(filter.Pagination.Offset()).Limit(filter.Pagination.PageSize)
	}

	arr, err := query.Order(
		ent.Desc(task.FieldUpdatedAt, task.FieldCreatedAt),
	).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return WarnPushEntArrToBiz(arr), total, nil
}

// CreateLog 批量查询
func (r *warnPushRepo) CreateLog(ctx context.Context, pushId uint64, log *biz.WarnPushLog) error {
	log.PushId = pushId
	entLog := WarnPushLogBizToEnt(log)
	return r.data.db.WarnPushLog(ctx).Create().SetWarnPushLog(entLog).Exec(ctx)
}

// ListLog 批量查询
func (r *warnPushRepo) ListLog(ctx context.Context, filter *biz.PushLogListFilter) ([]*biz.WarnPushLog, int, error) {
	query := r.data.db.WarnPushLog(ctx).Query()
	if filter != nil {
		if len(filter.PushIds) > 0 {
			query = query.Where(warnpushlog.PushIDIn(filter.PushIds...))
		}
		if filter.Status != "" {
			query = query.Where(warnpushlog.StatusEQ(filter.Status))
		}
		if filter.ParamContain != "" {
			query = query.Where(warnpushlog.ParamContains(filter.ParamContain))
		}
		if filter.BeginTime != nil {
			query = query.Where(warnpushlog.CreatedAtGTE(*filter.BeginTime))
		}
		if filter.EndTime != nil {
			query = query.Where(warnpushlog.CreatedAtLTE(*filter.EndTime))
		}
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 分页
	if filter.Pagination != nil {
		query = query.Clone().Offset(filter.Pagination.Offset()).Limit(filter.Pagination.PageSize)
	}

	arr, err := query.Clone().Order(
		ent.Desc(warnpushlog.FieldCreatedAt),
	).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return WarnPushLogEntArrToBiz(arr), total, nil
}

func (r *warnPushRepo) buildQueryByFilter(query *ent.WarnPushQuery, filter *biz.PushListFilter) *ent.WarnPushQuery {
	// 模糊查询条件
	if filter.NameContain != "" {
		query = query.Where(warnpush.NameContains(filter.NameContain))
	}

	if filter.Status != biz.WarnPushStatusUnknown {
		query = query.Where(warnpush.StatusEQ(filter.Status))
	}

	return query
}

func (r *warnPushRepo) query(ctx context.Context) *ent.WarnPushQuery {
	return r.data.db.WarnPush(ctx).Query()
}

// WarnPushLogEntToBiz 将WarnPushLog从ent层转化为biz层结构
func WarnPushLogEntToBiz(wpl *ent.WarnPushLog) *biz.WarnPushLog {
	log := &biz.WarnPushLog{
		Id:            wpl.ID,
		Param:         wpl.Param,
		Remark:        wpl.Remark,
		PushId:        wpl.PushID,
		Status:        wpl.Status,
		CreatedAt:     wpl.CreatedAt,
		UpdatedAt:     wpl.UpdatedAt,
		TenantID:      wpl.TenantID,
		AccessOrgList: wpl.AccessOrgList,
	}
	if wpl.Edges.Push != nil {
		log.Push = WarnPushEntToBiz(wpl.Edges.Push)
	}

	return log
}

// WarnPushLogBizToEnt 将WarnPushLog从biz层转化为ent层结构
func WarnPushLogBizToEnt(wpl *biz.WarnPushLog) *ent.WarnPushLog {
	log := &ent.WarnPushLog{
		CreatedAt:     wpl.CreatedAt,
		UpdatedAt:     wpl.UpdatedAt,
		Param:         wpl.Param,
		Remark:        wpl.Remark,
		Status:        wpl.Status,
		PushID:        wpl.PushId,
		TenantID:      wpl.TenantID,
		AccessOrgList: wpl.AccessOrgList,
	}

	return log
}

// WarnPushLogEntArrToBiz 将WarnPushLog数组从ent层转化为biz层结构
func WarnPushLogEntArrToBiz(arr []*ent.WarnPushLog) []*biz.WarnPushLog {
	return convert.ToArr[ent.WarnPushLog, biz.WarnPushLog](WarnPushLogEntToBiz, arr)
}
