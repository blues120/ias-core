// Code generated by ent, DO NOT EDIT.

package equipattr

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldDeletedAt, v))
}

// AttrKey applies equality check predicate on the "attr_key" field. It's identical to AttrKeyEQ.
func AttrKey(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldAttrKey, v))
}

// AttrValue applies equality check predicate on the "attr_value" field. It's identical to AttrValueEQ.
func AttrValue(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldAttrValue, v))
}

// Extend applies equality check predicate on the "extend" field. It's identical to ExtendEQ.
func Extend(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldExtend, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotNull(FieldDeletedAt))
}

// AttrKeyEQ applies the EQ predicate on the "attr_key" field.
func AttrKeyEQ(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldAttrKey, v))
}

// AttrKeyNEQ applies the NEQ predicate on the "attr_key" field.
func AttrKeyNEQ(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldAttrKey, v))
}

// AttrKeyIn applies the In predicate on the "attr_key" field.
func AttrKeyIn(vs ...string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldAttrKey, vs...))
}

// AttrKeyNotIn applies the NotIn predicate on the "attr_key" field.
func AttrKeyNotIn(vs ...string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldAttrKey, vs...))
}

// AttrKeyGT applies the GT predicate on the "attr_key" field.
func AttrKeyGT(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldAttrKey, v))
}

// AttrKeyGTE applies the GTE predicate on the "attr_key" field.
func AttrKeyGTE(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldAttrKey, v))
}

// AttrKeyLT applies the LT predicate on the "attr_key" field.
func AttrKeyLT(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldAttrKey, v))
}

// AttrKeyLTE applies the LTE predicate on the "attr_key" field.
func AttrKeyLTE(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldAttrKey, v))
}

// AttrKeyContains applies the Contains predicate on the "attr_key" field.
func AttrKeyContains(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldContains(FieldAttrKey, v))
}

// AttrKeyHasPrefix applies the HasPrefix predicate on the "attr_key" field.
func AttrKeyHasPrefix(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldHasPrefix(FieldAttrKey, v))
}

// AttrKeyHasSuffix applies the HasSuffix predicate on the "attr_key" field.
func AttrKeyHasSuffix(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldHasSuffix(FieldAttrKey, v))
}

// AttrKeyEqualFold applies the EqualFold predicate on the "attr_key" field.
func AttrKeyEqualFold(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEqualFold(FieldAttrKey, v))
}

// AttrKeyContainsFold applies the ContainsFold predicate on the "attr_key" field.
func AttrKeyContainsFold(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldContainsFold(FieldAttrKey, v))
}

// AttrValueEQ applies the EQ predicate on the "attr_value" field.
func AttrValueEQ(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldAttrValue, v))
}

// AttrValueNEQ applies the NEQ predicate on the "attr_value" field.
func AttrValueNEQ(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldAttrValue, v))
}

// AttrValueIn applies the In predicate on the "attr_value" field.
func AttrValueIn(vs ...string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldAttrValue, vs...))
}

// AttrValueNotIn applies the NotIn predicate on the "attr_value" field.
func AttrValueNotIn(vs ...string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldAttrValue, vs...))
}

// AttrValueGT applies the GT predicate on the "attr_value" field.
func AttrValueGT(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldAttrValue, v))
}

// AttrValueGTE applies the GTE predicate on the "attr_value" field.
func AttrValueGTE(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldAttrValue, v))
}

// AttrValueLT applies the LT predicate on the "attr_value" field.
func AttrValueLT(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldAttrValue, v))
}

// AttrValueLTE applies the LTE predicate on the "attr_value" field.
func AttrValueLTE(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldAttrValue, v))
}

// AttrValueContains applies the Contains predicate on the "attr_value" field.
func AttrValueContains(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldContains(FieldAttrValue, v))
}

// AttrValueHasPrefix applies the HasPrefix predicate on the "attr_value" field.
func AttrValueHasPrefix(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldHasPrefix(FieldAttrValue, v))
}

// AttrValueHasSuffix applies the HasSuffix predicate on the "attr_value" field.
func AttrValueHasSuffix(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldHasSuffix(FieldAttrValue, v))
}

// AttrValueEqualFold applies the EqualFold predicate on the "attr_value" field.
func AttrValueEqualFold(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEqualFold(FieldAttrValue, v))
}

// AttrValueContainsFold applies the ContainsFold predicate on the "attr_value" field.
func AttrValueContainsFold(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldContainsFold(FieldAttrValue, v))
}

// ExtendEQ applies the EQ predicate on the "extend" field.
func ExtendEQ(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEQ(FieldExtend, v))
}

// ExtendNEQ applies the NEQ predicate on the "extend" field.
func ExtendNEQ(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNEQ(FieldExtend, v))
}

// ExtendIn applies the In predicate on the "extend" field.
func ExtendIn(vs ...string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldIn(FieldExtend, vs...))
}

// ExtendNotIn applies the NotIn predicate on the "extend" field.
func ExtendNotIn(vs ...string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldNotIn(FieldExtend, vs...))
}

// ExtendGT applies the GT predicate on the "extend" field.
func ExtendGT(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGT(FieldExtend, v))
}

// ExtendGTE applies the GTE predicate on the "extend" field.
func ExtendGTE(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldGTE(FieldExtend, v))
}

// ExtendLT applies the LT predicate on the "extend" field.
func ExtendLT(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLT(FieldExtend, v))
}

// ExtendLTE applies the LTE predicate on the "extend" field.
func ExtendLTE(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldLTE(FieldExtend, v))
}

// ExtendContains applies the Contains predicate on the "extend" field.
func ExtendContains(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldContains(FieldExtend, v))
}

// ExtendHasPrefix applies the HasPrefix predicate on the "extend" field.
func ExtendHasPrefix(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldHasPrefix(FieldExtend, v))
}

// ExtendHasSuffix applies the HasSuffix predicate on the "extend" field.
func ExtendHasSuffix(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldHasSuffix(FieldExtend, v))
}

// ExtendEqualFold applies the EqualFold predicate on the "extend" field.
func ExtendEqualFold(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldEqualFold(FieldExtend, v))
}

// ExtendContainsFold applies the ContainsFold predicate on the "extend" field.
func ExtendContainsFold(v string) predicate.EquipAttr {
	return predicate.EquipAttr(sql.FieldContainsFold(FieldExtend, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EquipAttr) predicate.EquipAttr {
	return predicate.EquipAttr(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EquipAttr) predicate.EquipAttr {
	return predicate.EquipAttr(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EquipAttr) predicate.EquipAttr {
	return predicate.EquipAttr(sql.NotPredicates(p))
}