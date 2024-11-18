package data

import (
	"context"
	"github.com/blues120/ias-core/data/ent/eventsubscription"
	"github.com/blues120/ias-core/data/ent/task"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/convert"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent"
)

type eventSubsRepo struct {
	data *Data

	log *log.Helper
}

func NewEventSubsRepo(data *Data, logger log.Logger) biz.EventSubsRepo {
	return &eventSubsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func EventSubToEnt(a *biz.EventSubscription) *ent.EventSubscription {
	eventSub := &ent.EventSubscription{
		ID:         a.ID,
		BoxID:      a.BoxId,
		ChannelID:  a.ChannelId,
		Callback:   a.Callback,
		TemplateID: a.TemplateId,
		Status:     a.Status,
	}
	return eventSub
}

// EventSubEntArrToBiz 将EventSub数组从ent层转化为biz层结构
func EventSubEntArrToBiz(arr []*ent.EventSubscription) []*biz.EventSubscription {
	return convert.ToArr[ent.EventSubscription, biz.EventSubscription](EventSubToBiz, arr)
}

func EventSubToBiz(ca *ent.EventSubscription) *biz.EventSubscription {
	data := &biz.EventSubscription{
		ID:         ca.ID,
		BoxId:      ca.BoxID,
		ChannelId:  ca.ChannelID,
		TemplateId: ca.TemplateID,
		Callback:   ca.Callback,
		Status:     ca.Status,
	}
	return data
}

func (r *eventSubsRepo) Save(ctx context.Context, ca *biz.EventSubscription) (uint64, error) {
	data, err := r.data.db.EventSubscription(ctx).Create().SetEventSubscription(EventSubToEnt(ca)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *eventSubsRepo) Update(ctx context.Context, id uint64, ca *biz.EventSubscription) error {
	query := r.data.db.EventSubscription(ctx).
		UpdateOneID(id).
		SetBoxID(ca.BoxId).
		SetChannelID(ca.ChannelId).
		SetCallback(ca.Callback).
		SetTemplateID(ca.TemplateId).
		SetStatus(ca.Status)
	return query.Exec(ctx)
}

func (r *eventSubsRepo) updateStatus(ctx context.Context, id uint64, status biz.EventSubStatus) error {
	return r.data.db.EventSubscription(ctx).UpdateOneID(id).SetStatus(status).Exec(ctx)
}

func (r *eventSubsRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.EventSubscription(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *eventSubsRepo) Find(ctx context.Context, id uint64) (*biz.EventSubscription, error) {
	query := r.data.db.EventSubscription(ctx).Query().Where(eventsubscription.IDEQ(id))
	ca, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorSubscribeNotFound("订阅告警(id=%d)不存在", id)
		}
		return nil, err
	}
	return EventSubToBiz(ca), nil
}

func (r *eventSubsRepo) FindByFilter(ctx context.Context, filter *biz.EventSubsFindFilter) (*biz.EventSubscription, error) {
	query := r.data.db.EventSubscription(ctx).Query()
	if filter != nil {
		query = r.buildQueryByExtraParam(query, filter)
	}
	ca, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
			// return nil, errors.ErrorSubscribeNotFound("订阅告警不存在")
		}
		return nil, err
	}
	return EventSubToBiz(ca), nil
}

func (r *eventSubsRepo) buildQueryByExtraParam(query *ent.EventSubscriptionQuery, filter *biz.EventSubsFindFilter) *ent.EventSubscriptionQuery {
	if filter.ChannelId != "" {
		query = query.Where(eventsubscription.ChannelIDEQ(filter.ChannelId))
	}
	if filter.BoxId != "" {
		query = query.Where(eventsubscription.BoxIDEQ(filter.BoxId))
	}
	if filter.TemplateId != "" {
		query = query.Where(eventsubscription.TemplateIDEQ(filter.TemplateId))
	}
	if filter.Status != "" {
		query = query.Where(eventsubscription.StatusEQ(filter.Status))
	}
	return query
}

// List 批量查询
func (r *eventSubsRepo) List(ctx context.Context, filter *biz.EventSubsFilter) ([]*biz.EventSubscription, int, error) {
	query := r.query(ctx)
	if filter != nil {
		query = r.buildQueryByFilter(query, filter)
	}

	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	arr, err := r.buildQueryPagination(query, filter).Clone().Order(
		ent.Asc(task.FieldStatus),
		ent.Desc(task.FieldUpdatedAt, task.FieldCreatedAt),
	).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return EventSubEntArrToBiz(arr), total, nil
}

func (r *eventSubsRepo) buildQueryByFilter(query *ent.EventSubscriptionQuery, filter *biz.EventSubsFilter) *ent.EventSubscriptionQuery {
	// 模糊查询条件
	if len(filter.ChannelIds) > 0 {
		query = query.Where(eventsubscription.ChannelIDIn(filter.ChannelIds...))
	}
	if filter.ChannelId != "" {
		query = query.Where(eventsubscription.ChannelIDEQ(filter.ChannelId))
	}
	if filter.TemplateId != "" {
		query = query.Where(eventsubscription.TemplateIDEQ(filter.TemplateId))
	}
	if filter.BoxId != "" {
		query = query.Where(eventsubscription.BoxIDEQ(filter.BoxId))
	}
	if filter.Status != "" {
		query = query.Where(eventsubscription.StatusEQ(filter.Status))
	}
	return query
}

// 分页
func (r *eventSubsRepo) buildQueryPagination(query *ent.EventSubscriptionQuery, filter *biz.EventSubsFilter) *ent.EventSubscriptionQuery {
	if filter == nil || filter.Pagination == nil {
		return query
	}

	pagination := filter.Pagination
	return query.Offset(pagination.Offset()).Limit(pagination.PageSize)
}

func (r *eventSubsRepo) query(ctx context.Context) *ent.EventSubscriptionQuery {
	return r.data.db.EventSubscription(ctx).Query()
}

func (r *eventSubsRepo) CheckExists(ctx context.Context, req *biz.CheckEventSubsExistsRequest) (bool, error) {
	return true, nil
}

// Disablestatus 将推送状态置为disable
func (r *eventSubsRepo) BatchUpdatestatus(ctx context.Context, filter *biz.EventSubsFilter) error {
	_, err := r.data.db.EventSubscription(ctx).Update().
		Where(eventsubscription.BoxIDEQ(filter.BoxId)).
		Where(eventsubscription.ChannelIDIn(filter.ChannelIds...)).
		SetStatus(filter.Status).
		Save(ctx)
	return err
}
