// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"microservice/app/course/service/internal/data/ent/course"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseCreate is the builder for creating a Course entity.
type CourseCreate struct {
	config
	mutation *CourseMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (cc *CourseCreate) SetTitle(s string) *CourseCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CourseCreate) SetDescription(s string) *CourseCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CourseCreate) SetNillableDescription(s *string) *CourseCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetTeacherID sets the "teacher_id" field.
func (cc *CourseCreate) SetTeacherID(i int64) *CourseCreate {
	cc.mutation.SetTeacherID(i)
	return cc
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cc *CourseCreate) SetNillableTeacherID(i *int64) *CourseCreate {
	if i != nil {
		cc.SetTeacherID(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CourseCreate) SetID(i int64) *CourseCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the CourseMutation object of the builder.
func (cc *CourseCreate) Mutation() *CourseMutation {
	return cc.mutation
}

// Save creates the Course in the database.
func (cc *CourseCreate) Save(ctx context.Context) (*Course, error) {
	var (
		err  error
		node *Course
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CourseCreate) SaveX(ctx context.Context) *Course {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CourseCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CourseCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CourseCreate) check() error {
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	return nil
}

func (cc *CourseCreate) sqlSave(ctx context.Context) (*Course, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (cc *CourseCreate) createSpec() (*Course, *sqlgraph.CreateSpec) {
	var (
		_node = &Course{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: course.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: course.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := cc.mutation.TeacherID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: course.FieldTeacherID,
		})
		_node.TeacherID = value
	}
	return _node, _spec
}

// CourseCreateBulk is the builder for creating many Course entities in bulk.
type CourseCreateBulk struct {
	config
	builders []*CourseCreate
}

// Save creates the Course entities in the database.
func (ccb *CourseCreateBulk) Save(ctx context.Context) ([]*Course, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Course, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CourseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CourseCreateBulk) SaveX(ctx context.Context) []*Course {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CourseCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CourseCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
