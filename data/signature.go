package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/signature"
)

type signatureRepo struct {
	data *Data
	log  *log.Helper
}

func NewSignatureRepo(data *Data, logger log.Logger) biz.SignatureRepo {
	return &signatureRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *signatureRepo) Save(ctx context.Context, si *biz.Signature) (uint64, error) {
	data, err := r.data.db.Signature(ctx).Create().SetSignature(r.bizToEnt(si)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *signatureRepo) FindByCondition(ctx context.Context, boxId string, appId string) (*biz.Signature, error) {
	if boxId == "" && appId == "" {
		return nil, fmt.Errorf("boxId and appId can not be empty")
	}
	query := r.data.db.Signature(ctx).Query()
	if boxId != "" {
		query = query.Where(signature.BoxIDEQ(boxId))
	}
	if appId != "" {
		query = query.Where(signature.AppIDEQ(appId))
	}
	data, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return SignatureEntToBiz(data), nil
}

func (r *signatureRepo) bizToEnt(si *biz.Signature) *ent.Signature {
	return &ent.Signature{
		ID:        si.Id,
		CreatedAt: si.CreatedAt,
		UpdatedAt: si.UpdatedAt,
		BoxID:     si.BoxId,
		AppID:     si.AppId,
		AppSecret: si.AppSecret,
	}
}

func SignatureEntToBiz(ent *ent.Signature) *biz.Signature {
	return &biz.Signature{
		Id:        ent.ID,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
		BoxId:     ent.BoxID,
		AppId:     ent.AppID,
		AppSecret: ent.AppSecret,
	}
}
