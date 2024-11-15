package biz

import (
	"context"
	"io"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/blues120/ias-kit/oss"
)

type OssUsecase struct {
	ossRepo oss.Oss

	log *log.Helper
}

func NewOssUsecase(repo oss.Oss, logger log.Logger) *OssUsecase {
	return &OssUsecase{ossRepo: repo, log: log.NewHelper(logger)}
}

// Upload 上传文件
func (uc *OssUsecase) Upload(ctx context.Context, key string, reader io.Reader) error {
	return uc.ossRepo.Upload(ctx, key, reader)
}

// Download 下载文件
func (uc *OssUsecase) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	return uc.ossRepo.Download(ctx, key)
}

// GenerateTemporaryUrl 生成临时访问链接
// expire 链接过期时间，最长有效期7天
func (uc *OssUsecase) GenerateTemporaryUrl(ctx context.Context, key string, expire time.Duration) (string, error) {
	return uc.ossRepo.GenerateTemporaryUrl(ctx, key, expire)
}

// GeneratePermanentUrl 生成永久访问链接
func (uc *OssUsecase) GeneratePermanentUrl(ctx context.Context, key string) (string, error) {
	return uc.ossRepo.GeneratePermanentUrl(ctx, key)
}

// 初始化分片上传，用于大文件分片上传
func (uc *OssUsecase) CreateMultipartUpload(ctx context.Context, key string) (uploadId string, err error) {
	return uc.ossRepo.CreateMultipartUpload(ctx, key)
}

// 上传分片
func (uc *OssUsecase) UploadPart(ctx context.Context, key, uploadId string, partNumber int64, reader io.ReadSeeker) (etag string, err error) {
	return uc.ossRepo.UploadPart(ctx, key, uploadId, partNumber, reader)
}

// 终止分片上传
func (uc *OssUsecase) AbortMultipartUpload(ctx context.Context, key, uploadId string) error {
	return uc.ossRepo.AbortMultipartUpload(ctx, key, uploadId)
}

// 完成分片上传
func (uc *OssUsecase) CompleteMultipartUpload(ctx context.Context, key, uploadId string, partsNum int64) (etag string, err error) {
	return uc.ossRepo.CompleteMultipartUpload(ctx, key, uploadId, partsNum)
}

// 列举已上传分片
func (uc *OssUsecase) ListParts(ctx context.Context, key, uploadId string, maxParts int64) (parts []*oss.CompletedPart, err error) {
	return uc.ossRepo.ListParts(ctx, key, uploadId, maxParts)
}
