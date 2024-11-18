package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/ent/inform"
)

type informRepo struct {
	data *Data

	log *log.Helper
}

func NewInformRepo(data *Data, logger log.Logger) biz.InformRepo {
	return &informRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// InformEntToBiz 将ent中的Inform模型转换为biz中的Inform模型
func InformEntToBiz(a *ent.Inform) *biz.Inform {
	Inform := &biz.Inform{
		ID:             a.ID,
		AlarmName:      a.AlarmName,
		AlarmType:      a.AlarmType,
		SignName:       a.SignName,
		NotifyTemplate: a.NotifyTemplate,
		TemplateCode:   a.TemplateCode,
		PhoneNumbers:   a.PhoneNumbers,
		NotifySwitch:   a.NotifySwitch,
		TaskID:         a.TaskID,
		TaskName:       a.TaskName,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
		TenantID:       a.TenantID,
		AccessOrgList:  a.AccessOrgList,
	}
	return Inform
}

// InformBizToEnt 将biz中的Inform模型转换为ent中的Inform模型
func InformBizToEnt(a *biz.Inform) *ent.Inform {
	Inform := &ent.Inform{
		ID:             a.ID,
		AlarmName:      a.AlarmName,
		AlarmType:      a.AlarmType,
		SignName:       a.SignName,
		NotifyTemplate: a.NotifyTemplate,
		TemplateCode:   a.TemplateCode,
		PhoneNumbers:   a.PhoneNumbers,
		NotifySwitch:   a.NotifySwitch,
		TaskID:         a.TaskID,
		TaskName:       a.TaskName,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
		TenantID:       a.TenantID,
		AccessOrgList:  a.AccessOrgList,
	}
	return Inform
}

func (r *informRepo) List(ctx context.Context, filter *biz.InformFilter) ([]*biz.Inform, int, error) {
	query := r.data.db.Inform(ctx).Query()

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
	informs, err := query.All(ctx)
	if err != nil {
		return nil, 0, err
	}
	// 转换为业务层对象
	list := make([]*biz.Inform, len(informs))
	for i, inform := range informs {
		list[i] = InformEntToBiz(inform)
	}
	return list, total, nil
}

func (r *informRepo) Save(ctx context.Context, Inform *biz.Inform) (uint64, error) {
	data, err := r.data.db.Inform(ctx).Create().SetInform(InformBizToEnt(Inform)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *informRepo) Find(ctx context.Context, id uint64) (*biz.Inform, error) {
	foundInform, err := r.data.db.Inform(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return InformEntToBiz(foundInform), nil
}

func (r *informRepo) FindByTaskID(ctx context.Context, taskID uint64) (*biz.Inform, error) {
	foundInform, err := r.data.db.Inform(ctx).Query().Where(inform.TaskID(taskID)).First(ctx)
	if err != nil {
		return nil, err
	}
	return InformEntToBiz(foundInform), nil
}

func (r *informRepo) Update(ctx context.Context, Inform *biz.Inform) error {
	return r.data.db.Inform(ctx).UpdateOneID(Inform.ID).SetInform(InformBizToEnt(Inform)).Exec(ctx)
}

func (r *informRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.Inform(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *informRepo) DeleteByTaskID(ctx context.Context, taskID uint64) (int, error) {
	return r.data.db.Inform(ctx).Delete().Where(inform.TaskIDEQ(taskID)).Exec(ctx)
}
