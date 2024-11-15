// Code generated by ent, DO NOT EDIT.

package activeinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// ProcessID applies equality check predicate on the "process_id" field. It's identical to ProcessIDEQ.
func ProcessID(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldProcessID, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldStartTime, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldResult, v))
}

// Msg applies equality check predicate on the "msg" field. It's identical to MsgEQ.
func Msg(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldMsg, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotNull(FieldDeletedAt))
}

// ProcessIDEQ applies the EQ predicate on the "process_id" field.
func ProcessIDEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldProcessID, v))
}

// ProcessIDNEQ applies the NEQ predicate on the "process_id" field.
func ProcessIDNEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldProcessID, v))
}

// ProcessIDIn applies the In predicate on the "process_id" field.
func ProcessIDIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldProcessID, vs...))
}

// ProcessIDNotIn applies the NotIn predicate on the "process_id" field.
func ProcessIDNotIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldProcessID, vs...))
}

// ProcessIDGT applies the GT predicate on the "process_id" field.
func ProcessIDGT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldProcessID, v))
}

// ProcessIDGTE applies the GTE predicate on the "process_id" field.
func ProcessIDGTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldProcessID, v))
}

// ProcessIDLT applies the LT predicate on the "process_id" field.
func ProcessIDLT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldProcessID, v))
}

// ProcessIDLTE applies the LTE predicate on the "process_id" field.
func ProcessIDLTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldProcessID, v))
}

// ProcessIDContains applies the Contains predicate on the "process_id" field.
func ProcessIDContains(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContains(FieldProcessID, v))
}

// ProcessIDHasPrefix applies the HasPrefix predicate on the "process_id" field.
func ProcessIDHasPrefix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasPrefix(FieldProcessID, v))
}

// ProcessIDHasSuffix applies the HasSuffix predicate on the "process_id" field.
func ProcessIDHasSuffix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasSuffix(FieldProcessID, v))
}

// ProcessIDEqualFold applies the EqualFold predicate on the "process_id" field.
func ProcessIDEqualFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEqualFold(FieldProcessID, v))
}

// ProcessIDContainsFold applies the ContainsFold predicate on the "process_id" field.
func ProcessIDContainsFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContainsFold(FieldProcessID, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeContains applies the Contains predicate on the "start_time" field.
func StartTimeContains(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContains(FieldStartTime, v))
}

// StartTimeHasPrefix applies the HasPrefix predicate on the "start_time" field.
func StartTimeHasPrefix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasPrefix(FieldStartTime, v))
}

// StartTimeHasSuffix applies the HasSuffix predicate on the "start_time" field.
func StartTimeHasSuffix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasSuffix(FieldStartTime, v))
}

// StartTimeEqualFold applies the EqualFold predicate on the "start_time" field.
func StartTimeEqualFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEqualFold(FieldStartTime, v))
}

// StartTimeContainsFold applies the ContainsFold predicate on the "start_time" field.
func StartTimeContainsFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContainsFold(FieldStartTime, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldResult, v))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldResult, v))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldResult, vs...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldResult, vs...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldResult, v))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldResult, v))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldResult, v))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldResult, v))
}

// ResultContains applies the Contains predicate on the "result" field.
func ResultContains(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContains(FieldResult, v))
}

// ResultHasPrefix applies the HasPrefix predicate on the "result" field.
func ResultHasPrefix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasPrefix(FieldResult, v))
}

// ResultHasSuffix applies the HasSuffix predicate on the "result" field.
func ResultHasSuffix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasSuffix(FieldResult, v))
}

// ResultEqualFold applies the EqualFold predicate on the "result" field.
func ResultEqualFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEqualFold(FieldResult, v))
}

// ResultContainsFold applies the ContainsFold predicate on the "result" field.
func ResultContainsFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContainsFold(FieldResult, v))
}

// MsgEQ applies the EQ predicate on the "msg" field.
func MsgEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEQ(FieldMsg, v))
}

// MsgNEQ applies the NEQ predicate on the "msg" field.
func MsgNEQ(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNEQ(FieldMsg, v))
}

// MsgIn applies the In predicate on the "msg" field.
func MsgIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldIn(FieldMsg, vs...))
}

// MsgNotIn applies the NotIn predicate on the "msg" field.
func MsgNotIn(vs ...string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldNotIn(FieldMsg, vs...))
}

// MsgGT applies the GT predicate on the "msg" field.
func MsgGT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGT(FieldMsg, v))
}

// MsgGTE applies the GTE predicate on the "msg" field.
func MsgGTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldGTE(FieldMsg, v))
}

// MsgLT applies the LT predicate on the "msg" field.
func MsgLT(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLT(FieldMsg, v))
}

// MsgLTE applies the LTE predicate on the "msg" field.
func MsgLTE(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldLTE(FieldMsg, v))
}

// MsgContains applies the Contains predicate on the "msg" field.
func MsgContains(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContains(FieldMsg, v))
}

// MsgHasPrefix applies the HasPrefix predicate on the "msg" field.
func MsgHasPrefix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasPrefix(FieldMsg, v))
}

// MsgHasSuffix applies the HasSuffix predicate on the "msg" field.
func MsgHasSuffix(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldHasSuffix(FieldMsg, v))
}

// MsgEqualFold applies the EqualFold predicate on the "msg" field.
func MsgEqualFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldEqualFold(FieldMsg, v))
}

// MsgContainsFold applies the ContainsFold predicate on the "msg" field.
func MsgContainsFold(v string) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.FieldContainsFold(FieldMsg, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ActiveInfo) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ActiveInfo) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ActiveInfo) predicate.ActiveInfo {
	return predicate.ActiveInfo(sql.NotPredicates(p))
}
