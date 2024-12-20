// Code generated by ent, DO NOT EDIT.

package warnpushlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldDeletedAt, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldTenantID, v))
}

// AccessOrgList applies equality check predicate on the "access_org_list" field. It's identical to AccessOrgListEQ.
func AccessOrgList(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldAccessOrgList, v))
}

// PushID applies equality check predicate on the "push_id" field. It's identical to PushIDEQ.
func PushID(v uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldPushID, v))
}

// Param applies equality check predicate on the "param" field. It's identical to ParamEQ.
func Param(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldParam, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotNull(FieldDeletedAt))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContains(FieldTenantID, v))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasPrefix(FieldTenantID, v))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasSuffix(FieldTenantID, v))
}

// TenantIDIsNil applies the IsNil predicate on the "tenant_id" field.
func TenantIDIsNil() predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIsNull(FieldTenantID))
}

// TenantIDNotNil applies the NotNil predicate on the "tenant_id" field.
func TenantIDNotNil() predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotNull(FieldTenantID))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEqualFold(FieldTenantID, v))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContainsFold(FieldTenantID, v))
}

// AccessOrgListEQ applies the EQ predicate on the "access_org_list" field.
func AccessOrgListEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldAccessOrgList, v))
}

// AccessOrgListNEQ applies the NEQ predicate on the "access_org_list" field.
func AccessOrgListNEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldAccessOrgList, v))
}

// AccessOrgListIn applies the In predicate on the "access_org_list" field.
func AccessOrgListIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldAccessOrgList, vs...))
}

// AccessOrgListNotIn applies the NotIn predicate on the "access_org_list" field.
func AccessOrgListNotIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldAccessOrgList, vs...))
}

// AccessOrgListGT applies the GT predicate on the "access_org_list" field.
func AccessOrgListGT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldAccessOrgList, v))
}

// AccessOrgListGTE applies the GTE predicate on the "access_org_list" field.
func AccessOrgListGTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldAccessOrgList, v))
}

// AccessOrgListLT applies the LT predicate on the "access_org_list" field.
func AccessOrgListLT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldAccessOrgList, v))
}

// AccessOrgListLTE applies the LTE predicate on the "access_org_list" field.
func AccessOrgListLTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldAccessOrgList, v))
}

// AccessOrgListContains applies the Contains predicate on the "access_org_list" field.
func AccessOrgListContains(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContains(FieldAccessOrgList, v))
}

// AccessOrgListHasPrefix applies the HasPrefix predicate on the "access_org_list" field.
func AccessOrgListHasPrefix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasPrefix(FieldAccessOrgList, v))
}

// AccessOrgListHasSuffix applies the HasSuffix predicate on the "access_org_list" field.
func AccessOrgListHasSuffix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasSuffix(FieldAccessOrgList, v))
}

// AccessOrgListIsNil applies the IsNil predicate on the "access_org_list" field.
func AccessOrgListIsNil() predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIsNull(FieldAccessOrgList))
}

// AccessOrgListNotNil applies the NotNil predicate on the "access_org_list" field.
func AccessOrgListNotNil() predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotNull(FieldAccessOrgList))
}

// AccessOrgListEqualFold applies the EqualFold predicate on the "access_org_list" field.
func AccessOrgListEqualFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEqualFold(FieldAccessOrgList, v))
}

// AccessOrgListContainsFold applies the ContainsFold predicate on the "access_org_list" field.
func AccessOrgListContainsFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContainsFold(FieldAccessOrgList, v))
}

// PushIDEQ applies the EQ predicate on the "push_id" field.
func PushIDEQ(v uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldPushID, v))
}

// PushIDNEQ applies the NEQ predicate on the "push_id" field.
func PushIDNEQ(v uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldPushID, v))
}

// PushIDIn applies the In predicate on the "push_id" field.
func PushIDIn(vs ...uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldPushID, vs...))
}

// PushIDNotIn applies the NotIn predicate on the "push_id" field.
func PushIDNotIn(vs ...uint64) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldPushID, vs...))
}

// ParamEQ applies the EQ predicate on the "param" field.
func ParamEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldParam, v))
}

// ParamNEQ applies the NEQ predicate on the "param" field.
func ParamNEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldParam, v))
}

// ParamIn applies the In predicate on the "param" field.
func ParamIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldParam, vs...))
}

// ParamNotIn applies the NotIn predicate on the "param" field.
func ParamNotIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldParam, vs...))
}

// ParamGT applies the GT predicate on the "param" field.
func ParamGT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldParam, v))
}

// ParamGTE applies the GTE predicate on the "param" field.
func ParamGTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldParam, v))
}

// ParamLT applies the LT predicate on the "param" field.
func ParamLT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldParam, v))
}

// ParamLTE applies the LTE predicate on the "param" field.
func ParamLTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldParam, v))
}

// ParamContains applies the Contains predicate on the "param" field.
func ParamContains(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContains(FieldParam, v))
}

// ParamHasPrefix applies the HasPrefix predicate on the "param" field.
func ParamHasPrefix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasPrefix(FieldParam, v))
}

// ParamHasSuffix applies the HasSuffix predicate on the "param" field.
func ParamHasSuffix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasSuffix(FieldParam, v))
}

// ParamEqualFold applies the EqualFold predicate on the "param" field.
func ParamEqualFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEqualFold(FieldParam, v))
}

// ParamContainsFold applies the ContainsFold predicate on the "param" field.
func ParamContainsFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContainsFold(FieldParam, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.FieldContainsFold(FieldRemark, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v biz.WarnPushLogStatus) predicate.WarnPushLog {
	vc := v
	return predicate.WarnPushLog(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v biz.WarnPushLogStatus) predicate.WarnPushLog {
	vc := v
	return predicate.WarnPushLog(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...biz.WarnPushLogStatus) predicate.WarnPushLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPushLog(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...biz.WarnPushLogStatus) predicate.WarnPushLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPushLog(sql.FieldNotIn(FieldStatus, v...))
}

// HasPush applies the HasEdge predicate on the "push" edge.
func HasPush() predicate.WarnPushLog {
	return predicate.WarnPushLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PushTable, PushColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPushWith applies the HasEdge predicate on the "push" edge with a given conditions (other predicates).
func HasPushWith(preds ...predicate.WarnPush) predicate.WarnPushLog {
	return predicate.WarnPushLog(func(s *sql.Selector) {
		step := newPushStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WarnPushLog) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WarnPushLog) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WarnPushLog) predicate.WarnPushLog {
	return predicate.WarnPushLog(sql.NotPredicates(p))
}
