package biz

import (
	"context"
	"errors"

	"github.com/blues120/ias-core/pkg/pagination"
)

type FileUploadStatus string

// Values provides list valid values for Enum.
func (FileUploadStatus) Values() []string {
	return []string{
		string(FileUploadStatusUnknown),
		string(FileUploadStatusUploading),
		string(FileUploadStatusSuccess),
		string(FileUploadStatusFailed),
	}
}

const (
	FileUploadStatusUnknown   FileUploadStatus = "unknown"   // 未知
	FileUploadStatusUploading FileUploadStatus = "uploading" // 上传中
	FileUploadStatusSuccess   FileUploadStatus = "success"   // 上传成功
	FileUploadStatusFailed    FileUploadStatus = "failed"    // 上传失败
)

type FileUpload struct {
	ID               uint64           // id
	Provider         string           // 供应商
	FileName         string           // 文件名
	Md5              string           // 文件 md5
	TotalBytes       uint64           // 文件大小字节数
	Etag             string           // 文件 etag
	Key              string           // 文件 oss key
	Url              string           // 文件 oss url
	UploadID         string           // 文件上传 id
	CreateTime       string           // 上传时间
	Status           FileUploadStatus // 文件上传状态
	AlgoGroupID      uint64           // 算法组id
	Type             string           // 类型,docker image或者文件
	Meta             string           // 安装meta信息
	Platform         string           // 设备平台
	DeviceModel      string           // 设备型号
	AlgoGroupName    string           // 算法组名称
	AlgoGroupVersion string           // 算法组版本
}

type FileUploadFilter struct {
	FileNameEq string // 文件名称
	Md5Eq      string // md5
	KeyEq      string // 文件 oss key
	ProviderEq string // 供应商

	/*
		范围查询条件
	*/
	IDIn []uint64 // 包含文件 ID
	/*
		分页
	*/
	Pagination *pagination.Pagination
}

type FileUploadRepo interface {
	List(ctx context.Context, filter *FileUploadFilter) ([]*FileUpload, int, error)
	Save(ctx context.Context, fileUpload *FileUpload) (uint64, error)
	Update(ctx context.Context, fileUpload *FileUpload) error
	Delete(ctx context.Context, id uint64) error
	FindByCondition(ctx context.Context, fileUpload *FileUpload) (bool, error)
	Find(ctx context.Context, id uint64) (*FileUpload, error)
}

type FileUploadUsecase struct {
	repo FileUploadRepo
}

func NewFileUploadUsecase(repo FileUploadRepo) *FileUploadUsecase {
	return &FileUploadUsecase{repo: repo}
}

// List 根据条件查询
func (uc *FileUploadUsecase) List(ctx context.Context, filter *FileUploadFilter) ([]*FileUpload, int, error) {
	return uc.repo.List(ctx, filter)
}

// Create 校验是否重复并入库
func (uc *FileUploadUsecase) Create(ctx context.Context, fileUpload *FileUpload) (uint64, error) {
	// md5 重复校验
	isExist, err := uc.repo.FindByCondition(ctx, fileUpload)
	if err != nil {
		return 0, err
	}
	if isExist {
		return 0, errors.New("file already exists")
	}

	return uc.repo.Save(ctx, fileUpload)
}

// Update 校验是否重名并update库
func (uc *FileUploadUsecase) Update(ctx context.Context, id uint64, fileUpload *FileUpload) error {
	// 文件ID是否存在
	update, err := uc.repo.Find(ctx, id)
	if err != nil {
		return err
	}

	// Md5 重复校验
	if update.Md5 != fileUpload.Md5 {
		isMd5Exist := &FileUpload{Md5: fileUpload.Md5}
		isExist, err := uc.repo.FindByCondition(ctx, isMd5Exist)
		if err != nil {
			return err
		}
		if isExist {
			return errors.New("file already exists")
		}
	}

	return uc.repo.Update(ctx, fileUpload)
}

// Delete 根据id删除
func (uc *FileUploadUsecase) Delete(ctx context.Context, id uint64) error {
	// 文件ID是否存在
	fileUpload, err := uc.repo.Find(ctx, id)
	if err != nil {
		return err
	}

	// 上传中的文件禁止删除
	if fileUpload.Status == FileUploadStatusUploading {
		return errors.New("file is uploading, please wait")
	}

	return uc.repo.Delete(ctx, id)
}

func (uc *FileUploadUsecase) Find(ctx context.Context, id uint64) (*FileUpload, error) {
	return uc.repo.Find(ctx, id)
}
