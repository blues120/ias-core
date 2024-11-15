package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/area"
	"strconv"
)

type areaRepo struct {
	data *Data
	log  *log.Helper
}

func NewAreaRepo(data *Data, logger log.Logger) biz.AreaRepo {
	return &areaRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *areaRepo) List(ctx context.Context, level uint32, pid int32) ([]*biz.Area, error) {
	query := r.data.db.Area(ctx).Query().Where(area.Pid(int64(pid))).Where(area.Level(uint64(level)))

	all, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为业务层对象
	list := make([]*biz.Area, 0)
	for _, ar := range all {
		list = append(list, AreaEntToBiz(ar))
	}
	return list, nil
}

func (r *areaRepo) QueryNamesByIds(ctx context.Context, id []uint64) ([]string, error) {
	query := r.data.db.Area(ctx).Query().Where(area.IDIn(id...))
	all, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为业务层对象
	list := make([]string, 0)
	for _, ar := range all {
		list = append(list, ar.Name)
	}
	return list, nil
}

func (r *areaRepo) QueryByIds(ctx context.Context, id []uint64) ([]*biz.Area, error) {
	query := r.data.db.Area(ctx).Query().Where(area.IDIn(id...))
	all, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为业务层对象
	list := make([]*biz.Area, 0)
	for _, ar := range all {
		list = append(list, AreaEntToBiz(ar))
	}
	return list, nil
}

func (r *areaRepo) QueryNameIdPidMap(ctx context.Context) (map[string]string, map[string]string, error) {
	areaNameIdMap := make(map[string]string)
	areaNamePidMap := make(map[string]string)
	query := r.data.db.Area(ctx).Query()
	all, err := query.All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for _, a := range all {
		if idstr, ok := areaNameIdMap[a.Name]; ok {
			idstr += fmt.Sprintf(",%v", a.ID)
			areaNameIdMap[a.Name] = idstr

			pidStr := areaNamePidMap[a.Name]
			pidStr += fmt.Sprintf(",%v", a.Pid)
			areaNamePidMap[a.Name] = pidStr
		} else {
			areaNameIdMap[a.Name] = strconv.Itoa(int(a.ID))
			areaNamePidMap[a.Name] = strconv.Itoa(int(a.Pid))
		}
	}
	return areaNameIdMap, areaNamePidMap, nil
}

// AreaEntToBiz 将ent中的area模型转换为biz中的area模型
func AreaEntToBiz(a *ent.Area) *biz.Area {
	return &biz.Area{
		ID:    a.ID,
		Name:  a.Name,
		Level: a.Level,
		Pid:   a.Pid,
	}
}
