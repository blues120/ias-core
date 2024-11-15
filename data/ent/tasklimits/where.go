// Code generated by ent, DO NOT EDIT.

package tasklimits

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldDeletedAt, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldModel, v))
}

// MaxCameraNum applies equality check predicate on the "maxCameraNum" field. It's identical to MaxCameraNumEQ.
func MaxCameraNum(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldMaxCameraNum, v))
}

// AlgoNum applies equality check predicate on the "algoNum" field. It's identical to AlgoNumEQ.
func AlgoNum(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldAlgoNum, v))
}

// MaxSubTaskNum applies equality check predicate on the "maxSubTaskNum" field. It's identical to MaxSubTaskNumEQ.
func MaxSubTaskNum(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldMaxSubTaskNum, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotNull(FieldDeletedAt))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldHasSuffix(FieldModel, v))
}

// ModelIsNil applies the IsNil predicate on the "model" field.
func ModelIsNil() predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIsNull(FieldModel))
}

// ModelNotNil applies the NotNil predicate on the "model" field.
func ModelNotNil() predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotNull(FieldModel))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldContainsFold(FieldModel, v))
}

// MaxCameraNumEQ applies the EQ predicate on the "maxCameraNum" field.
func MaxCameraNumEQ(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldMaxCameraNum, v))
}

// MaxCameraNumNEQ applies the NEQ predicate on the "maxCameraNum" field.
func MaxCameraNumNEQ(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldMaxCameraNum, v))
}

// MaxCameraNumIn applies the In predicate on the "maxCameraNum" field.
func MaxCameraNumIn(vs ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldMaxCameraNum, vs...))
}

// MaxCameraNumNotIn applies the NotIn predicate on the "maxCameraNum" field.
func MaxCameraNumNotIn(vs ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldMaxCameraNum, vs...))
}

// MaxCameraNumGT applies the GT predicate on the "maxCameraNum" field.
func MaxCameraNumGT(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldMaxCameraNum, v))
}

// MaxCameraNumGTE applies the GTE predicate on the "maxCameraNum" field.
func MaxCameraNumGTE(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldMaxCameraNum, v))
}

// MaxCameraNumLT applies the LT predicate on the "maxCameraNum" field.
func MaxCameraNumLT(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldMaxCameraNum, v))
}

// MaxCameraNumLTE applies the LTE predicate on the "maxCameraNum" field.
func MaxCameraNumLTE(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldMaxCameraNum, v))
}

// AlgoNumEQ applies the EQ predicate on the "algoNum" field.
func AlgoNumEQ(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldAlgoNum, v))
}

// AlgoNumNEQ applies the NEQ predicate on the "algoNum" field.
func AlgoNumNEQ(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldAlgoNum, v))
}

// AlgoNumIn applies the In predicate on the "algoNum" field.
func AlgoNumIn(vs ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldAlgoNum, vs...))
}

// AlgoNumNotIn applies the NotIn predicate on the "algoNum" field.
func AlgoNumNotIn(vs ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldAlgoNum, vs...))
}

// AlgoNumGT applies the GT predicate on the "algoNum" field.
func AlgoNumGT(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldAlgoNum, v))
}

// AlgoNumGTE applies the GTE predicate on the "algoNum" field.
func AlgoNumGTE(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldAlgoNum, v))
}

// AlgoNumLT applies the LT predicate on the "algoNum" field.
func AlgoNumLT(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldAlgoNum, v))
}

// AlgoNumLTE applies the LTE predicate on the "algoNum" field.
func AlgoNumLTE(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldAlgoNum, v))
}

// MaxSubTaskNumEQ applies the EQ predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumEQ(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldEQ(FieldMaxSubTaskNum, v))
}

// MaxSubTaskNumNEQ applies the NEQ predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumNEQ(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNEQ(FieldMaxSubTaskNum, v))
}

// MaxSubTaskNumIn applies the In predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumIn(vs ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldIn(FieldMaxSubTaskNum, vs...))
}

// MaxSubTaskNumNotIn applies the NotIn predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumNotIn(vs ...uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldNotIn(FieldMaxSubTaskNum, vs...))
}

// MaxSubTaskNumGT applies the GT predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumGT(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGT(FieldMaxSubTaskNum, v))
}

// MaxSubTaskNumGTE applies the GTE predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumGTE(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldGTE(FieldMaxSubTaskNum, v))
}

// MaxSubTaskNumLT applies the LT predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumLT(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLT(FieldMaxSubTaskNum, v))
}

// MaxSubTaskNumLTE applies the LTE predicate on the "maxSubTaskNum" field.
func MaxSubTaskNumLTE(v uint64) predicate.TaskLimits {
	return predicate.TaskLimits(sql.FieldLTE(FieldMaxSubTaskNum, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskLimits) predicate.TaskLimits {
	return predicate.TaskLimits(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskLimits) predicate.TaskLimits {
	return predicate.TaskLimits(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TaskLimits) predicate.TaskLimits {
	return predicate.TaskLimits(sql.NotPredicates(p))
}
