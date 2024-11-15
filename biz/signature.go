package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type SignatureRepo interface {
	// Save 创建签名
	Save(ctx context.Context, si *Signature) (uint64, error)

	// FindByCondition 查询签名
	FindByCondition(ctx context.Context, boxId string, appId string) (*Signature, error)
}

type Signature struct {
	Id        uint64    // id
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	BoxId     string    // 盒子Id
	AppId     string    // AppId
	AppSecret string    // App秘钥
}

type SignatureUsecase struct {
	signatureRepo SignatureRepo

	log *log.Helper
}

func NewSignatureUsecase(repo SignatureRepo, logger log.Logger) *SignatureUsecase {
	return &SignatureUsecase{signatureRepo: repo, log: log.NewHelper(logger)}
}

// Create 创建签名
func (uc *SignatureUsecase) Create(ctx context.Context, si *Signature) (uint64, error) {
	return uc.signatureRepo.Save(ctx, si)
}

// FindByCondition 查询签名
func (uc *SignatureUsecase) FindByCondition(ctx context.Context, boxId string, appId string) (*Signature, error) {
	return uc.signatureRepo.FindByCondition(ctx, boxId, appId)
}
