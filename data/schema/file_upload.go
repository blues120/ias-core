package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/mixin"
)

type FileUpload struct {
	ent.Schema
}

func (FileUpload) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("provider").Comment("文件提供者，比如算法文件来自 ctyun_ias/sophgo_park/sophgo_city"),
		field.String("file_name").Comment("文件名"),
		field.String("md5").Comment("文件 md5"),
		field.Uint64("total_bytes").Comment("文件大小字节数，前端显示用"),
		field.String("etag").Optional().Comment("文件 etag,不一定等于 md5"),
		field.String("key").Comment("oss 存储的 key"),
		field.String("upload_id").Optional().Comment("oss 上传的标识，用于断点续传"),
		field.String("status").Comment("上传状态"),
		field.String("type").Comment("类型,docker image或者文件"),
		field.Uint64("algo_group_id").Comment("算法组id"),
		field.Text("meta").Comment("安装meta信息"),
	}
}

// Annotations of the FileUpload
func (FileUpload) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "file_upload"},
	}
}

// Mixin of the FileUpload
func (FileUpload) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}
