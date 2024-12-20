// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/blues120/ias-core/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldParentID, v))
}

// IamRoleID applies equality check predicate on the "iam_role_id" field. It's identical to IamRoleIDEQ.
func IamRoleID(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldIamRoleID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldName, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...uint32) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...uint32) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v uint32) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldParentID, v))
}

// IamRoleIDEQ applies the EQ predicate on the "iam_role_id" field.
func IamRoleIDEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldIamRoleID, v))
}

// IamRoleIDNEQ applies the NEQ predicate on the "iam_role_id" field.
func IamRoleIDNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldIamRoleID, v))
}

// IamRoleIDIn applies the In predicate on the "iam_role_id" field.
func IamRoleIDIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldIamRoleID, vs...))
}

// IamRoleIDNotIn applies the NotIn predicate on the "iam_role_id" field.
func IamRoleIDNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldIamRoleID, vs...))
}

// IamRoleIDGT applies the GT predicate on the "iam_role_id" field.
func IamRoleIDGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldIamRoleID, v))
}

// IamRoleIDGTE applies the GTE predicate on the "iam_role_id" field.
func IamRoleIDGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldIamRoleID, v))
}

// IamRoleIDLT applies the LT predicate on the "iam_role_id" field.
func IamRoleIDLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldIamRoleID, v))
}

// IamRoleIDLTE applies the LTE predicate on the "iam_role_id" field.
func IamRoleIDLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldIamRoleID, v))
}

// IamRoleIDContains applies the Contains predicate on the "iam_role_id" field.
func IamRoleIDContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldIamRoleID, v))
}

// IamRoleIDHasPrefix applies the HasPrefix predicate on the "iam_role_id" field.
func IamRoleIDHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldIamRoleID, v))
}

// IamRoleIDHasSuffix applies the HasSuffix predicate on the "iam_role_id" field.
func IamRoleIDHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldIamRoleID, v))
}

// IamRoleIDEqualFold applies the EqualFold predicate on the "iam_role_id" field.
func IamRoleIDEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldIamRoleID, v))
}

// IamRoleIDContainsFold applies the ContainsFold predicate on the "iam_role_id" field.
func IamRoleIDContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldIamRoleID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.NotPredicates(p))
}
