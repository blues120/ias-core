package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/activeinfo"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/equipattr"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/convert"
)

type systemRepo struct {
	data *Data
	log  *log.Helper
}

// NewSystemRepo
func NewSystemRepo(data *Data, logger log.Logger) biz.SystemRepo {
	return &systemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// List 批量查询
func (r *systemRepo) ListEquipAttr(ctx context.Context, filter *biz.EquipAttrtFilter) ([]*biz.EquipAttr, error) {

	query := r.query(ctx)
	if filter != nil {
		if filter.AttrKey != "" {
			query = query.Where(equipattr.AttrKeyEQ(filter.AttrKey))
		}
	}

	arr, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return EquipAttrEntArrToBiz(arr), nil

}

func (r *systemRepo) query(ctx context.Context) *ent.EquipAttrQuery {
	return r.data.db.EquipAttr(ctx).Query()
}

func EquipAttrEntArrToBiz(arr []*ent.EquipAttr) []*biz.EquipAttr {
	return convert.ToArr[ent.EquipAttr, biz.EquipAttr](EquipAttrEntToBiz, arr)
}

func EquipAttrEntToBiz(ent *ent.EquipAttr) *biz.EquipAttr {
	equipAttr := &biz.EquipAttr{
		Id:        ent.ID,
		AttrKey:   ent.AttrKey,
		AttrValue: ent.AttrValue,
		Extend:    ent.Extend,
	}
	return equipAttr
}

func (r *systemRepo) InsertEquipAttr(ctx context.Context, attrKey string, attrValue string) (bool, error) {

	if _, err := r.data.db.EquipAttr(ctx).Create().
		SetAttrKey(attrKey).
		SetAttrValue(attrValue).
		SetExtend("").
		Save(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (r *systemRepo) UpdateEquipAttr(ctx context.Context, attrKey string, attrValue string) (bool, error) {
	query := r.query(ctx).Where(equipattr.AttrKeyEQ(attrKey))
	attrInfo, err := query.First(ctx)
	if err != nil {
		return false, err
	}

	if err = r.data.db.EquipAttr(ctx).UpdateOneID(attrInfo.ID).SetAttrValue(attrValue).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (r *systemRepo) GetEquipAttr(ctx context.Context, attrKey string) (string, error) {
	query := r.query(ctx).Where(equipattr.AttrKeyEQ(attrKey))
	attrInfo, err := query.First(ctx)
	if err != nil {
		return "", err
	}

	return attrInfo.AttrValue, nil
}

func (r *systemRepo) InsertActiveInfo(ctx context.Context, process_id string) (bool, error) {
	if _, err := r.data.db.ActiveInfo(ctx).Create().
		SetProcessID(process_id).
		SetStartTime(time.Now().String()).
		SetResult("activating").
		SetMsg("激活中").
		Save(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (r *systemRepo) DeleteEquipAttr(ctx context.Context, attrKey string) (int, error) {

	return r.data.db.EquipAttr(ctx).Delete().Where(equipattr.AttrKeyEQ(attrKey)).Exec(ctx)
}

func (r *systemRepo) UpdateActiveInfo(ctx context.Context, process_id string, result string, msg string) (bool, error) {
	query := r.data.db.ActiveInfo(ctx).Query().Where(activeinfo.ProcessIDEQ(process_id))
	attrInfo, err := query.First(ctx)
	if err != nil {
		return false, err // todo 未找到的错误
	}
	r.data.db.ActiveInfo(ctx).UpdateOneID(attrInfo.ID).SetResult(result).SetMsg(msg).Exec(ctx)
	return true, nil
}

func (r *systemRepo) GetActiveInfo(ctx context.Context, process_id string) (*biz.ActiveInfo, error) {
	query := r.data.db.ActiveInfo(ctx).Query().Where(activeinfo.ProcessIDEQ(process_id))
	attrInfo, err := query.First(ctx)
	if err != nil {
		return nil, err // todo 未找到的错误
	}
	return AttrInfoEntToBiz(attrInfo), nil
}

func (r *systemRepo) GetVersion(ctx context.Context) (string, error) {
	data, err := r.data.db.Setting(ctx).Query().First(ctx)
	if err != nil {
		return "", err
	}
	return data.Version, nil
}

func (r *systemRepo) UpdateVersion(ctx context.Context, newVersion string) error {
	return r.data.db.Setting(ctx).Update().SetVersion(newVersion).Exec(ctx)
}

func AttrInfoEntToBiz(ent *ent.ActiveInfo) *biz.ActiveInfo {
	activeInfo := &biz.ActiveInfo{
		Id:        ent.ID,
		ProcessId: ent.ProcessID,
		StartTime: ent.StartTime,
		Result:    ent.Result,
		Msg:       ent.Msg,
	}
	return activeInfo
}
