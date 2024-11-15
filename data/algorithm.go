package data

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/algorithm"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

type algoRepo struct {
	data *Data

	log *log.Helper
}

func NewAlgoRepo(data *Data, logger log.Logger) biz.AlgoRepo {
	return &algoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// AlgoEntToBiz 将ent中的Algorithm模型转换为biz中的Algorithm模型
func AlgoEntToBiz(a *ent.Algorithm) *biz.Algorithm {
	algo := &biz.Algorithm{
		ID:               a.ID,
		Name:             a.Name,
		Type:             a.Type,
		Desc:             a.Description,
		Version:          a.Version,
		AlarmType:        a.AlarmType,
		AlarmName:        a.AlarmName,
		Notify:           a.Notify,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
		DrawType:         a.DrawType,
		BaseType:         a.BaseType,
		AppName:          a.AppName,
		Image:            a.Image,
		LabelMap:         a.LabelMap,
		Target:           a.Target,
		AlgoNameEn:       a.AlgoNameEn,
		AlgoGroupID:      a.AlgoGroupID,
		AlgoGroupName:    a.AlgoGroupName,
		AlgoGroupVersion: a.AlgoGroupVersion,
		Config:           a.Config,
		Provider:         a.Provider,
		AlgoID:           a.AlgoID,
		Available:        a.Available,
		Platform:         a.Platform,
		DeviceModel:      a.DeviceModel,
		IsGroupType:      a.IsGroupType,
		Prefix:           a.Prefix,
	}
	extendStr, _ := json.Marshal(a.Extend)
	if a.Type == "frame" {
		algo.Detail = string(extendStr)
		// 从extend json拆出resultType
		var aie biz.AlgoFrameExtend
		_ = json.Unmarshal(extendStr, &aie)
		algo.Notify = aie.Notify
	} else {
		// 从extend json拆分fileName和configName
		var ale biz.AlgoStreamExtend
		_ = json.Unmarshal(extendStr, &ale)
		algo.FileName = ale.FileName
		algo.ConfigName = ale.ConfigName
	}
	return algo
}

func AlgoBizToEnt(a *biz.Algorithm) *ent.Algorithm {
	var extent map[string]interface{}
	if a.Extend != "" {
		err := json.Unmarshal([]byte(a.Extend), &extent)
		if err != nil {
			log.Errorf("Error unmarshaling JSON: %v", err)
		}
	}
	algo := &ent.Algorithm{
		ID:               a.ID,
		Name:             a.Name,
		Type:             a.Type,
		Description:      a.Desc,
		Version:          a.Version,
		AlarmType:        a.AlarmType,
		AlarmName:        a.AlarmName,
		Notify:           a.Notify,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
		DrawType:         a.DrawType,
		BaseType:         a.BaseType,
		Extend:           extent,
		AppName:          a.AppName,
		Image:            a.Image,
		LabelMap:         a.LabelMap,
		Target:           a.Target,
		AlgoNameEn:       a.AlgoNameEn,
		AlgoGroupID:      a.AlgoGroupID,
		AlgoGroupName:    a.AlgoGroupName,
		AlgoGroupVersion: a.AlgoGroupVersion,
		Config:           a.Config,
		Provider:         a.Provider,
		AlgoID:           a.AlgoID,
		Available:        a.Available,
		Platform:         a.Platform,
		DeviceModel:      a.DeviceModel,
		IsGroupType:      a.IsGroupType,
		Prefix:           a.Prefix,
	}
	return algo
}

func (r *algoRepo) List(ctx context.Context, filter *biz.AlgoFilter) ([]*biz.Algorithm, int, error) {
	query := r.data.db.Algorithm(ctx).Query()

	// 组装过滤条件
	// ...
	if filter.NameEq != "" {
		query = query.Where(algorithm.NameEQ(filter.NameEq))
	}

	if filter.NameContains != "" {
		query = query.Where(algorithm.NameContains(filter.NameContains))
	}

	if filter.AvailableEq != nil {
		query = query.Where(algorithm.AvailableEQ(*filter.AvailableEq))
	}

	if filter.AlgoGroupIDEq != 0 {
		query = query.Where(algorithm.AlgoGroupIDEQ(filter.AlgoGroupIDEq))
	}

	if filter.AlgoGroupNameEq != "" {
		query = query.Where(algorithm.AlgoGroupNameEQ(filter.AlgoGroupNameEq))
	}

	if filter.AlgoGroupVersionEq != "" {
		query = query.Where(algorithm.AlgoGroupVersionEQ(filter.AlgoGroupVersionEq))
	}

	// 英文名称过滤
	if filter.AlgoNameEn != "" {
		query = query.Where(algorithm.AlgoNameEnEQ(filter.AlgoNameEn))
	}

	if len(filter.IDIn) > 0 {
		query = query.Where(algorithm.IDIn(filter.IDIn...))
	}

	if filter.ProviderEq != "" {
		query = query.Where(algorithm.ProviderEQ(filter.ProviderEq))
	}

	if filter.AlgoGroupNameContains != "" {
		query = query.Where(algorithm.AlgoGroupNameContains(filter.AlgoGroupNameContains))
	}

	if filter.PlatformEq != "" {
		query = query.Where(algorithm.PlatformEQ(filter.PlatformEq))
	}

	if filter.DeviceModelEq != "" {
		query = query.Where(algorithm.DeviceModelEQ(filter.DeviceModelEq))
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
	list := make([]*biz.Algorithm, len(algorithms))
	for i, algo := range algorithms {
		list[i] = AlgoEntToBiz(algo)
	}

	return list, total, nil
}

func (r *algoRepo) Save(ctx context.Context, algo *biz.Algorithm) (uint64, error) {
	data, err := r.data.db.Algorithm(ctx).Create().SetAlgorithm(AlgoBizToEnt(algo)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *algoRepo) FindByCondition(ctx context.Context, algo *biz.Algorithm) (bool, error) {
	query := r.data.db.Algorithm(ctx).Query()

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

func (r *algoRepo) buildAlgorithmPredicates(algo *biz.Algorithm) []predicate.Algorithm {
	var predicates []predicate.Algorithm
	value := reflect.ValueOf(*algo)
	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		v := value.Field(i)

		switch field.Name {
		case "ID":
			if id := v.Uint(); id != 0 {
				predicates = append(predicates, algorithm.IDEQ(id))
			}
		case "Name":
			if name := v.String(); name != "" {
				predicates = append(predicates, algorithm.NameEQ(name))
			}
		case "Type":
			if typ := v.String(); typ != "" {
				predicates = append(predicates, algorithm.TypeEQ(typ))
			}
		case "AlarmName":
			if alarmName := v.String(); alarmName != "" {
				predicates = append(predicates, algorithm.AlarmNameEQ(alarmName))
			}
		case "Provider":
			if typ := v.String(); typ != "" {
				predicates = append(predicates, algorithm.ProviderEQ(typ))
			}
		// 根据你的需求，继续添加其他字段
		case "AlgoGroupID":
			if agid := v.Uint(); agid != 0 {
				predicates = append(predicates, algorithm.AlgoGroupID(uint(agid)))
			}
		}
	}

	return predicates
}

func (r *algoRepo) Find(ctx context.Context, id uint64) (*biz.Algorithm, error) {
	foundAlgo, err := r.data.db.Algorithm(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return AlgoEntToBiz(foundAlgo), nil
}

func (r *algoRepo) FindOneByAlgoNameEn(ctx context.Context, algoNameEn string) (*biz.Algorithm, error) {
	algo, err := r.data.db.Algorithm(ctx).Query().Where(algorithm.AlgoNameEn(algoNameEn)).First(ctx)
	if err != nil {
		return nil, err
	}
	return AlgoEntToBiz(algo), nil
}

func (r *algoRepo) FindAlgorithms(ctx context.Context, filter *biz.AlgoFilter) ([]*biz.Algorithm, error) {
	query := r.data.db.Algorithm(ctx).Query()

	// 组装过滤条件
	// ...
	if filter.NameEq != "" {
		query = query.Where(algorithm.NameEQ(filter.NameEq))
	}

	if filter.ProviderEq != "" {
		query = query.Where(algorithm.ProviderEQ(filter.ProviderEq))
	}

	algorithm, err := query.All(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
		return nil, err
	}
	// 转换为业务层对象
	algolist := make([]*biz.Algorithm, 0)

	for _, v := range algorithm {
		algo := AlgoEntToBiz(v)
		algolist = append(algolist, algo)
	}
	return algolist, nil
}
func (r *algoRepo) Update(ctx context.Context, algo *biz.Algorithm) error {
	return r.data.db.Algorithm(ctx).UpdateOneID(algo.ID).SetAlgorithm(AlgoBizToEnt(algo)).Exec(ctx)
}

func (r *algoRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.Algorithm(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *algoRepo) FindTasksByAlgorithmID(ctx context.Context, id uint64) ([]*biz.Task, error) {
	algo, err := r.data.db.Algorithm(ctx).
		Query().
		Where(algorithm.ID(id)).
		WithTasks().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return TaskEntArrToBiz(algo.Edges.Tasks), nil
}

func (r *algoRepo) ResetAvailableAlgo(ctx context.Context, IDList []uint64) error {
	// 首先将所有的算法都置为不可用
	_, err := r.data.db.Algorithm(ctx).Update().SetAvailable(0).Save(context.Background())
	if err != nil {
		return err
	}

	// 如果不为空，然后将标记的置为可用
	if len(IDList) > 0 {
		_, err = r.data.db.Algorithm(ctx).Update().Where(algorithm.IDIn(IDList...)).SetAvailable(1).Save(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *algoRepo) Count(ctx context.Context) (int, error) {
	return r.data.db.Algorithm(ctx).Query().Count(ctx)
}

func (r *algoRepo) UpdateAlgoGroupVersion(ctx context.Context, name, version string) error {
	return r.data.db.Algorithm(ctx).Update().Where(algorithm.AlgoGroupNameEQ(name)).SetAlgoGroupVersion(version).Exec(ctx)
}
