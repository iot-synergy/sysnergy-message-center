// Code generated by ent, DO NOT EDIT.

package emaillog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-message-center/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// Target applies equality check predicate on the "target" field. It's identical to TargetEQ.
func Target(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldTarget, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldSubject, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldContent, v))
}

// SendStatus applies equality check predicate on the "send_status" field. It's identical to SendStatusEQ.
func SendStatus(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldSendStatus, v))
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldProvider, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// TargetEQ applies the EQ predicate on the "target" field.
func TargetEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldTarget, v))
}

// TargetNEQ applies the NEQ predicate on the "target" field.
func TargetNEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldTarget, v))
}

// TargetIn applies the In predicate on the "target" field.
func TargetIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldTarget, vs...))
}

// TargetNotIn applies the NotIn predicate on the "target" field.
func TargetNotIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldTarget, vs...))
}

// TargetGT applies the GT predicate on the "target" field.
func TargetGT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldTarget, v))
}

// TargetGTE applies the GTE predicate on the "target" field.
func TargetGTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldTarget, v))
}

// TargetLT applies the LT predicate on the "target" field.
func TargetLT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldTarget, v))
}

// TargetLTE applies the LTE predicate on the "target" field.
func TargetLTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldTarget, v))
}

// TargetContains applies the Contains predicate on the "target" field.
func TargetContains(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContains(FieldTarget, v))
}

// TargetHasPrefix applies the HasPrefix predicate on the "target" field.
func TargetHasPrefix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasPrefix(FieldTarget, v))
}

// TargetHasSuffix applies the HasSuffix predicate on the "target" field.
func TargetHasSuffix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasSuffix(FieldTarget, v))
}

// TargetEqualFold applies the EqualFold predicate on the "target" field.
func TargetEqualFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEqualFold(FieldTarget, v))
}

// TargetContainsFold applies the ContainsFold predicate on the "target" field.
func TargetContainsFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContainsFold(FieldTarget, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContainsFold(FieldSubject, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContainsFold(FieldContent, v))
}

// SendStatusEQ applies the EQ predicate on the "send_status" field.
func SendStatusEQ(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldSendStatus, v))
}

// SendStatusNEQ applies the NEQ predicate on the "send_status" field.
func SendStatusNEQ(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldSendStatus, v))
}

// SendStatusIn applies the In predicate on the "send_status" field.
func SendStatusIn(vs ...uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldSendStatus, vs...))
}

// SendStatusNotIn applies the NotIn predicate on the "send_status" field.
func SendStatusNotIn(vs ...uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldSendStatus, vs...))
}

// SendStatusGT applies the GT predicate on the "send_status" field.
func SendStatusGT(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldSendStatus, v))
}

// SendStatusGTE applies the GTE predicate on the "send_status" field.
func SendStatusGTE(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldSendStatus, v))
}

// SendStatusLT applies the LT predicate on the "send_status" field.
func SendStatusLT(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldSendStatus, v))
}

// SendStatusLTE applies the LTE predicate on the "send_status" field.
func SendStatusLTE(v uint8) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldSendStatus, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldNotIn(FieldProvider, vs...))
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGT(FieldProvider, v))
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldGTE(FieldProvider, v))
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLT(FieldProvider, v))
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldLTE(FieldProvider, v))
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContains(FieldProvider, v))
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasPrefix(FieldProvider, v))
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldHasSuffix(FieldProvider, v))
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldEqualFold(FieldProvider, v))
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.EmailLog {
	return predicate.EmailLog(sql.FieldContainsFold(FieldProvider, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailLog) predicate.EmailLog {
	return predicate.EmailLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailLog) predicate.EmailLog {
	return predicate.EmailLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EmailLog) predicate.EmailLog {
	return predicate.EmailLog(sql.NotPredicates(p))
}
