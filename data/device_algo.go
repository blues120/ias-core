package data

import (
	"context"
	"errors"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/ent/devicealgo"
	"github.com/blues120/ias-core/data/ent/predicate"
)

type deviceAlgoRepo struct {
	data *Data

	log *log.Helper
}

func NewDeviceAlgoRepo(data *Data, logger log.Logger) biz.DeviceAlgoRepo {
	return &deviceAlgoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func DeviceAlgoEntToBiz(a *ent.DeviceAlgo) *biz.DeviceAlgo {
	algo := &biz.DeviceAlgo{
		ID:               a.ID,
		Name:             a.Name,
		DeviceId:         a.DeviceID,
		AlgoGroupID:      a.AlgoGroupID,
		AlgoGroupName:    a.AlgoGroupName,
		AlgoGroupVersion: a.AlgoGroupVersion,
		InstallTime:      a.InstallTime,
		Version:          a.Version,
	}
	return algo
}

func DeviceAlgoBizToEnt(a *biz.DeviceAlgo) *ent.DeviceAlgo {
	algo := &ent.DeviceAlgo{
		ID:               a.ID,
		Name:             a.Name,
		DeviceID:         a.DeviceId,
		AlgoGroupID:      a.AlgoGroupID,
		AlgoGroupName:    a.AlgoGroupName,
		AlgoGroupVersion: a.AlgoGroupVersion,
		InstallTime:      a.InstallTime,
		Version: 		  a.Version,
	}
	return algo
}

// 查询条件
func (r *deviceAlgoRepo) buildQueryByFilter(query *ent.DeviceAlgoQuery, filter *biz.DeviceAlgoFilter) *ent.DeviceAlgoQuery {
	// 精确查询条件
	if filter.DeviceId != 0 {
		query = query.Where(devicealgo.DeviceIDEQ(filter.DeviceId))
	}
	if filter.AlgoGroupId != 0 {
		query = query.Where(devicealgo.AlgoGroupIDEQ(filter.AlgoGroupId))
	}
	if len(filter.DeviceIds) != 0 {
		query = query.Where(devicealgo.DeviceIDIn(filter.DeviceIds...))
	}
	return query
}

func (r *deviceAlgoRepo) List(ctx context.Context, filter *biz.DeviceAlgoFilter) ([]*biz.DeviceAlgo, int, error) {
	query := r.data.db.DeviceAlgo(ctx).Query()

	// 组装查询条件
	if filter != nil {
		query = r.buildQueryByFilter(query, filter)
	}

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
	algorithms, err := query.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 转换为业务层对象
	list := make([]*biz.DeviceAlgo, len(algorithms))
	for i, algo := range algorithms {
		list[i] = DeviceAlgoEntToBiz(algo)
	}

	return list, total, nil
}

func (r *deviceAlgoRepo) Save(ctx context.Context, algo *biz.DeviceAlgo) (uint64, error) {
	data, err := r.data.db.DeviceAlgo(ctx).Create().SetDeviceAlgo(DeviceAlgoBizToEnt(algo)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *deviceAlgoRepo) Update(ctx context.Context, algo *biz.DeviceAlgo) error {
	return r.data.db.DeviceAlgo(ctx).UpdateOneID(algo.ID).SetDeviceAlgo(DeviceAlgoBizToEnt(algo)).Exec(ctx)
}

func (r *deviceAlgoRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.DeviceAlgo(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *deviceAlgoRepo) Count(ctx context.Context) (int, error) {
	return r.data.db.DeviceAlgo(ctx).Query().Count(ctx)
}

func (r *deviceAlgoRepo) FindByCondition(ctx context.Context, algo *biz.DeviceAlgo) (bool, error) {
	query := r.data.db.DeviceAlgo(ctx).Query()

	predicates := r.buildAlgorithmPredicates(algo)
	query = query.Where(predicates...)

	foundAlgo, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	if foundAlgo != nil {
		return true, nil
	}

	return false, errors.New("unexpected error")
}

func (r *deviceAlgoRepo) Find(ctx context.Context, id uint64) (*biz.DeviceAlgo, error) {
	foundAlgo, err := r.data.db.DeviceAlgo(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return DeviceAlgoEntToBiz(foundAlgo), nil
}

func (r *deviceAlgoRepo) buildAlgorithmPredicates(algo *biz.DeviceAlgo) []predicate.DeviceAlgo {
	var predicates []predicate.DeviceAlgo
	value := reflect.ValueOf(*algo)
	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		v := value.Field(i)

		switch field.Name {
		case "ID":
			if id := v.Uint(); id != 0 {
				predicates = append(predicates, devicealgo.IDEQ(id))
			}
		case "DeviceId":
			if deviceId := v.Uint(); deviceId != 0 {
				predicates = append(predicates, devicealgo.DeviceIDEQ(deviceId))
			}
		case "AlgoGroupID":
			if algoGroupId := v.Uint(); algoGroupId != 0 {
				predicates = append(predicates, devicealgo.AlgoGroupIDEQ(uint(algoGroupId)))
			}
		case "AlgoGroupName":
			if algoGroupName := v.String(); algoGroupName != "" {
				predicates = append(predicates, devicealgo.AlgoGroupNameEQ(algoGroupName))
			}
		case "Name":
			if name := v.String(); name != "" {
				predicates = append(predicates, devicealgo.NameEQ(name))
			}
		case "Version":
			if version := v.String(); version != "" {
				predicates = append(predicates, devicealgo.VersionEQ(version))
			}
		}
	}

	return predicates
}

func (r *deviceAlgoRepo) BatchUpdateDeviceAlgo(ctx context.Context, deviceId uint64, algoList []*biz.DeviceAlgo) error {
	err := r.DeleteByDeviceId(ctx, deviceId)
	if err != nil {
		return err
	}
	return r.BatchCreate(ctx, algoList)
}

// 批量添加
func (r *deviceAlgoRepo) BatchCreate(ctx context.Context, algoList []*biz.DeviceAlgo) error {
	if len(algoList) == 0 {
		return nil
	}
	deviceAlgo := r.data.db.DeviceAlgo(ctx)
	var bulk []*ent.DeviceAlgoCreate
	for _, algo := range algoList {
		bulk = append(bulk, deviceAlgo.Create().SetDeviceAlgo(DeviceAlgoBizToEnt(algo)))
	}
	_, err := deviceAlgo.CreateBulk(bulk...).Save(ctx)
	return err
}

func (r *deviceAlgoRepo) DeleteByDeviceId(ctx context.Context, deviceId uint64) error {
	_, err := r.data.db.DeviceAlgo(ctx).Delete().Where(devicealgo.DeviceIDEQ(deviceId)).Exec(ctx)
	return err
}
