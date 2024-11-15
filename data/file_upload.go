package data

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/fileupload"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
)

type fileUploadRepo struct {
	data *Data

	log *log.Helper
}

func NewFileUploadRepo(data *Data, logger log.Logger) biz.FileUploadRepo {
	return &fileUploadRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// FileUploadEntToBiz 将ent中的FileUpload模型转换为biz中的FileUpload模型
func FileUploadEntToBiz(a *ent.FileUpload) *biz.FileUpload {
	fileUpload := &biz.FileUpload{
		ID:          a.ID,
		Provider:    a.Provider,
		FileName:    a.FileName,
		Md5:         a.Md5,
		TotalBytes:  a.TotalBytes,
		Etag:        a.Etag,
		Key:         a.Key,
		UploadID:    a.UploadID,
		CreateTime:  a.CreatedAt.In(time.Local).Format(DateTimeFormat),
		Status:      biz.FileUploadStatus(a.Status),
		Type:        a.Type,
		AlgoGroupID: a.AlgoGroupID,
		Meta:        a.Meta,
	}
	return fileUpload
}

func FileUploadBizToEnt(a *biz.FileUpload) *ent.FileUpload {
	fileUpload := &ent.FileUpload{
		ID:          a.ID,
		Provider:    a.Provider,
		FileName:    a.FileName,
		Md5:         a.Md5,
		TotalBytes:  a.TotalBytes,
		Etag:        a.Etag,
		Key:         a.Key,
		UploadID:    a.UploadID,
		Status:      string(a.Status),
		Type:        a.Type,
		AlgoGroupID: a.AlgoGroupID,
		Meta:        a.Meta,
	}
	return fileUpload
}

func (r *fileUploadRepo) List(ctx context.Context, filter *biz.FileUploadFilter) ([]*biz.FileUpload, int, error) {
	query := r.data.db.FileUpload(ctx).Query()

	// 组装过滤条件
	// ...
	if filter.FileNameEq != "" {
		query = query.Where(fileupload.FileNameEQ(filter.FileNameEq))
	}

	if filter.Md5Eq != "" {
		query = query.Where(fileupload.Md5EQ(filter.Md5Eq))
	}

	if filter.KeyEq != "" {
		query = query.Where(fileupload.KeyEQ(filter.KeyEq))
	}

	if len(filter.IDIn) > 0 {
		query = query.Where(fileupload.IDIn(filter.IDIn...))
	}

	if filter.ProviderEq != "" {
		query = query.Where(fileupload.ProviderEQ(filter.ProviderEq))
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
	fileUploads, err := query.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 转换为业务层对象
	list := make([]*biz.FileUpload, len(fileUploads))
	for i, fu := range fileUploads {
		list[i] = FileUploadEntToBiz(fu)
	}

	return list, total, nil
}

func (r *fileUploadRepo) Save(ctx context.Context, fu *biz.FileUpload) (uint64, error) {
	data, err := r.data.db.FileUpload(ctx).Create().SetFileUpload(FileUploadBizToEnt(fu)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *fileUploadRepo) FindByCondition(ctx context.Context, fu *biz.FileUpload) (bool, error) {
	query := r.data.db.FileUpload(ctx).Query()

	predicates := r.buildFileUploadPredicates(fu)
	query = query.Where(predicates...)

	foundFile, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	if foundFile != nil {
		return true, nil
	}

	return false, errors.New("unexpected error")
}

func (r *fileUploadRepo) buildFileUploadPredicates(fu *biz.FileUpload) []predicate.FileUpload {
	var predicates []predicate.FileUpload
	value := reflect.ValueOf(*fu)
	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		v := value.Field(i)

		switch field.Name {
		case "ID":
			if id := v.Uint(); id != 0 {
				predicates = append(predicates, fileupload.IDEQ(id))
			}
		case "FileName":
			if name := v.String(); name != "" {
				predicates = append(predicates, fileupload.FileNameEQ(name))
			}
		case "Md5":
			if typ := v.String(); typ != "" {
				predicates = append(predicates, fileupload.Md5EQ(typ))
			}
		case "Key":
			if typ := v.String(); typ != "" {
				predicates = append(predicates, fileupload.KeyEQ(typ))
			}
		case "Provider":
			if typ := v.String(); typ != "" {
				predicates = append(predicates, fileupload.ProviderEQ(typ))
			}
			// 根据你的需求，继续添加其他字段
		}
	}

	return predicates
}

func (r *fileUploadRepo) Find(ctx context.Context, id uint64) (*biz.FileUpload, error) {
	foundFileUpload, err := r.data.db.FileUpload(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return FileUploadEntToBiz(foundFileUpload), nil
}

func (r *fileUploadRepo) Update(ctx context.Context, fu *biz.FileUpload) error {
	return r.data.db.FileUpload(ctx).UpdateOneID(fu.ID).SetFileUpload(FileUploadBizToEnt(fu)).Exec(ctx)
}

func (r *fileUploadRepo) Delete(ctx context.Context, id uint64) error {
	return r.data.db.FileUpload(ctx).DeleteOneID(id).Exec(ctx)
}
