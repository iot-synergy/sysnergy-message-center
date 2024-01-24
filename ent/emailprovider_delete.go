// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/synergy-message-center/ent/emailprovider"
	"github.com/iot-synergy/synergy-message-center/ent/predicate"
)

// EmailProviderDelete is the builder for deleting a EmailProvider entity.
type EmailProviderDelete struct {
	config
	hooks    []Hook
	mutation *EmailProviderMutation
}

// Where appends a list predicates to the EmailProviderDelete builder.
func (epd *EmailProviderDelete) Where(ps ...predicate.EmailProvider) *EmailProviderDelete {
	epd.mutation.Where(ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *EmailProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, epd.sqlExec, epd.mutation, epd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *EmailProviderDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *EmailProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emailprovider.Table, sqlgraph.NewFieldSpec(emailprovider.FieldID, field.TypeUint64))
	if ps := epd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, epd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	epd.mutation.done = true
	return affected, err
}

// EmailProviderDeleteOne is the builder for deleting a single EmailProvider entity.
type EmailProviderDeleteOne struct {
	epd *EmailProviderDelete
}

// Where appends a list predicates to the EmailProviderDelete builder.
func (epdo *EmailProviderDeleteOne) Where(ps ...predicate.EmailProvider) *EmailProviderDeleteOne {
	epdo.epd.mutation.Where(ps...)
	return epdo
}

// Exec executes the deletion query.
func (epdo *EmailProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emailprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *EmailProviderDeleteOne) ExecX(ctx context.Context) {
	if err := epdo.Exec(ctx); err != nil {
		panic(err)
	}
}
