package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var (
	_ ent.Mixin = (*CreatedAt)(nil)
	_ ent.Mixin = (*UpdatedAt)(nil)
	_ ent.Mixin = (*Time)(nil)
)

// CreatedAt adds created at time field.
type CreatedAt struct{ mixin.Schema }

// Fields of the CreatedAt mixin.
func (CreatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			Immutable().
			Comment("创建时间"),
	}
}

// UpdatedAt adds updated at time field.
type UpdatedAt struct{ mixin.Schema }

// Fields of the UpdatedAt mixin.
func (UpdatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}

// Time composes CreatedAt / UpdatedAt mixin.
type Time struct{ mixin.Schema }

// Fields of the time mixin.
func (Time) Fields() []ent.Field {
	return append(
		CreatedAt{}.Fields(),
		UpdatedAt{}.Fields()...,
	)
}
