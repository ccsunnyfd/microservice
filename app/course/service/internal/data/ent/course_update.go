// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"microservice/app/course/service/internal/data/ent/course"
	"microservice/app/course/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseUpdate is the builder for updating Course entities.
type CourseUpdate struct {
	config
	hooks    []Hook
	mutation *CourseMutation
}

// Where appends a list predicates to the CourseUpdate builder.
func (cu *CourseUpdate) Where(ps ...predicate.Course) *CourseUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *CourseUpdate) SetTitle(s string) *CourseUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CourseUpdate) SetDescription(s string) *CourseUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableDescription(s *string) *CourseUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *CourseUpdate) ClearDescription() *CourseUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetTeacherID sets the "teacher_id" field.
func (cu *CourseUpdate) SetTeacherID(i int64) *CourseUpdate {
	cu.mutation.ResetTeacherID()
	cu.mutation.SetTeacherID(i)
	return cu
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableTeacherID(i *int64) *CourseUpdate {
	if i != nil {
		cu.SetTeacherID(*i)
	}
	return cu
}

// AddTeacherID adds i to the "teacher_id" field.
func (cu *CourseUpdate) AddTeacherID(i int64) *CourseUpdate {
	cu.mutation.AddTeacherID(i)
	return cu
}

// ClearTeacherID clears the value of the "teacher_id" field.
func (cu *CourseUpdate) ClearTeacherID() *CourseUpdate {
	cu.mutation.ClearTeacherID()
	return cu
}

// Mutation returns the CourseMutation object of the builder.
func (cu *CourseUpdate) Mutation() *CourseMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CourseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CourseUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CourseUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CourseUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   course.Table,
			Columns: course.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: course.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldTitle,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldDescription,
		})
	}
	if cu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: course.FieldDescription,
		})
	}
	if value, ok := cu.mutation.TeacherID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: course.FieldTeacherID,
		})
	}
	if value, ok := cu.mutation.AddedTeacherID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: course.FieldTeacherID,
		})
	}
	if cu.mutation.TeacherIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: course.FieldTeacherID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CourseUpdateOne is the builder for updating a single Course entity.
type CourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseMutation
}

// SetTitle sets the "title" field.
func (cuo *CourseUpdateOne) SetTitle(s string) *CourseUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CourseUpdateOne) SetDescription(s string) *CourseUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableDescription(s *string) *CourseUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *CourseUpdateOne) ClearDescription() *CourseUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetTeacherID sets the "teacher_id" field.
func (cuo *CourseUpdateOne) SetTeacherID(i int64) *CourseUpdateOne {
	cuo.mutation.ResetTeacherID()
	cuo.mutation.SetTeacherID(i)
	return cuo
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableTeacherID(i *int64) *CourseUpdateOne {
	if i != nil {
		cuo.SetTeacherID(*i)
	}
	return cuo
}

// AddTeacherID adds i to the "teacher_id" field.
func (cuo *CourseUpdateOne) AddTeacherID(i int64) *CourseUpdateOne {
	cuo.mutation.AddTeacherID(i)
	return cuo
}

// ClearTeacherID clears the value of the "teacher_id" field.
func (cuo *CourseUpdateOne) ClearTeacherID() *CourseUpdateOne {
	cuo.mutation.ClearTeacherID()
	return cuo
}

// Mutation returns the CourseMutation object of the builder.
func (cuo *CourseUpdateOne) Mutation() *CourseMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CourseUpdateOne) Select(field string, fields ...string) *CourseUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Course entity.
func (cuo *CourseUpdateOne) Save(ctx context.Context) (*Course, error) {
	var (
		err  error
		node *Course
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CourseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CourseUpdateOne) SaveX(ctx context.Context) *Course {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CourseUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CourseUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CourseUpdateOne) sqlSave(ctx context.Context) (_node *Course, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   course.Table,
			Columns: course.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: course.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Course.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, course.FieldID)
		for _, f := range fields {
			if !course.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != course.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: course.FieldDescription,
		})
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: course.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.TeacherID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: course.FieldTeacherID,
		})
	}
	if value, ok := cuo.mutation.AddedTeacherID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: course.FieldTeacherID,
		})
	}
	if cuo.mutation.TeacherIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: course.FieldTeacherID,
		})
	}
	_node = &Course{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
