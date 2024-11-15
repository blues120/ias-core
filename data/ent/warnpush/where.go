// Code generated by ent, DO NOT EDIT.

package warnpush

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldDeletedAt, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldTenantID, v))
}

// AccessOrgList applies equality check predicate on the "access_org_list" field. It's identical to AccessOrgListEQ.
func AccessOrgList(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldAccessOrgList, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldURL, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldRemark, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotNull(FieldDeletedAt))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContains(FieldTenantID, v))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasPrefix(FieldTenantID, v))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasSuffix(FieldTenantID, v))
}

// TenantIDIsNil applies the IsNil predicate on the "tenant_id" field.
func TenantIDIsNil() predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIsNull(FieldTenantID))
}

// TenantIDNotNil applies the NotNil predicate on the "tenant_id" field.
func TenantIDNotNil() predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotNull(FieldTenantID))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEqualFold(FieldTenantID, v))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContainsFold(FieldTenantID, v))
}

// AccessOrgListEQ applies the EQ predicate on the "access_org_list" field.
func AccessOrgListEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldAccessOrgList, v))
}

// AccessOrgListNEQ applies the NEQ predicate on the "access_org_list" field.
func AccessOrgListNEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldAccessOrgList, v))
}

// AccessOrgListIn applies the In predicate on the "access_org_list" field.
func AccessOrgListIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldAccessOrgList, vs...))
}

// AccessOrgListNotIn applies the NotIn predicate on the "access_org_list" field.
func AccessOrgListNotIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldAccessOrgList, vs...))
}

// AccessOrgListGT applies the GT predicate on the "access_org_list" field.
func AccessOrgListGT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldAccessOrgList, v))
}

// AccessOrgListGTE applies the GTE predicate on the "access_org_list" field.
func AccessOrgListGTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldAccessOrgList, v))
}

// AccessOrgListLT applies the LT predicate on the "access_org_list" field.
func AccessOrgListLT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldAccessOrgList, v))
}

// AccessOrgListLTE applies the LTE predicate on the "access_org_list" field.
func AccessOrgListLTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldAccessOrgList, v))
}

// AccessOrgListContains applies the Contains predicate on the "access_org_list" field.
func AccessOrgListContains(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContains(FieldAccessOrgList, v))
}

// AccessOrgListHasPrefix applies the HasPrefix predicate on the "access_org_list" field.
func AccessOrgListHasPrefix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasPrefix(FieldAccessOrgList, v))
}

// AccessOrgListHasSuffix applies the HasSuffix predicate on the "access_org_list" field.
func AccessOrgListHasSuffix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasSuffix(FieldAccessOrgList, v))
}

// AccessOrgListIsNil applies the IsNil predicate on the "access_org_list" field.
func AccessOrgListIsNil() predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIsNull(FieldAccessOrgList))
}

// AccessOrgListNotNil applies the NotNil predicate on the "access_org_list" field.
func AccessOrgListNotNil() predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotNull(FieldAccessOrgList))
}

// AccessOrgListEqualFold applies the EqualFold predicate on the "access_org_list" field.
func AccessOrgListEqualFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEqualFold(FieldAccessOrgList, v))
}

// AccessOrgListContainsFold applies the ContainsFold predicate on the "access_org_list" field.
func AccessOrgListContainsFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContainsFold(FieldAccessOrgList, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v biz.WarnPushType) predicate.WarnPush {
	vc := v
	return predicate.WarnPush(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v biz.WarnPushType) predicate.WarnPush {
	vc := v
	return predicate.WarnPush(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...biz.WarnPushType) predicate.WarnPush {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPush(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...biz.WarnPushType) predicate.WarnPush {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPush(sql.FieldNotIn(FieldType, v...))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContainsFold(FieldURL, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.WarnPush {
	return predicate.WarnPush(sql.FieldContainsFold(FieldRemark, v))
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v biz.WarnPushMode) predicate.WarnPush {
	vc := v
	return predicate.WarnPush(sql.FieldEQ(FieldMode, vc))
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v biz.WarnPushMode) predicate.WarnPush {
	vc := v
	return predicate.WarnPush(sql.FieldNEQ(FieldMode, vc))
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...biz.WarnPushMode) predicate.WarnPush {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPush(sql.FieldIn(FieldMode, v...))
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...biz.WarnPushMode) predicate.WarnPush {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPush(sql.FieldNotIn(FieldMode, v...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v biz.WarnPushStatus) predicate.WarnPush {
	vc := v
	return predicate.WarnPush(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v biz.WarnPushStatus) predicate.WarnPush {
	vc := v
	return predicate.WarnPush(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...biz.WarnPushStatus) predicate.WarnPush {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPush(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...biz.WarnPushStatus) predicate.WarnPush {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WarnPush(sql.FieldNotIn(FieldStatus, v...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WarnPush) predicate.WarnPush {
	return predicate.WarnPush(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WarnPush) predicate.WarnPush {
	return predicate.WarnPush(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WarnPush) predicate.WarnPush {
	return predicate.WarnPush(sql.NotPredicates(p))
}
