// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"microservice/app/user/service/internal/data/ent/predicate"
	"microservice/app/user/service/internal/data/ent/user"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	name          *string
	password      *string
	real_name     *string
	mobile        *string
	email         *string
	is_teacher    *bool
	stars         *int32
	addstars      *int32
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int64) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetRealName sets the "real_name" field.
func (m *UserMutation) SetRealName(s string) {
	m.real_name = &s
}

// RealName returns the value of the "real_name" field in the mutation.
func (m *UserMutation) RealName() (r string, exists bool) {
	v := m.real_name
	if v == nil {
		return
	}
	return *v, true
}

// OldRealName returns the old "real_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRealName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRealName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRealName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRealName: %w", err)
	}
	return oldValue.RealName, nil
}

// ClearRealName clears the value of the "real_name" field.
func (m *UserMutation) ClearRealName() {
	m.real_name = nil
	m.clearedFields[user.FieldRealName] = struct{}{}
}

// RealNameCleared returns if the "real_name" field was cleared in this mutation.
func (m *UserMutation) RealNameCleared() bool {
	_, ok := m.clearedFields[user.FieldRealName]
	return ok
}

// ResetRealName resets all changes to the "real_name" field.
func (m *UserMutation) ResetRealName() {
	m.real_name = nil
	delete(m.clearedFields, user.FieldRealName)
}

// SetMobile sets the "mobile" field.
func (m *UserMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *UserMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ClearMobile clears the value of the "mobile" field.
func (m *UserMutation) ClearMobile() {
	m.mobile = nil
	m.clearedFields[user.FieldMobile] = struct{}{}
}

// MobileCleared returns if the "mobile" field was cleared in this mutation.
func (m *UserMutation) MobileCleared() bool {
	_, ok := m.clearedFields[user.FieldMobile]
	return ok
}

// ResetMobile resets all changes to the "mobile" field.
func (m *UserMutation) ResetMobile() {
	m.mobile = nil
	delete(m.clearedFields, user.FieldMobile)
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ClearEmail clears the value of the "email" field.
func (m *UserMutation) ClearEmail() {
	m.email = nil
	m.clearedFields[user.FieldEmail] = struct{}{}
}

// EmailCleared returns if the "email" field was cleared in this mutation.
func (m *UserMutation) EmailCleared() bool {
	_, ok := m.clearedFields[user.FieldEmail]
	return ok
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
	delete(m.clearedFields, user.FieldEmail)
}

// SetIsTeacher sets the "is_teacher" field.
func (m *UserMutation) SetIsTeacher(b bool) {
	m.is_teacher = &b
}

// IsTeacher returns the value of the "is_teacher" field in the mutation.
func (m *UserMutation) IsTeacher() (r bool, exists bool) {
	v := m.is_teacher
	if v == nil {
		return
	}
	return *v, true
}

// OldIsTeacher returns the old "is_teacher" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldIsTeacher(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldIsTeacher is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldIsTeacher requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsTeacher: %w", err)
	}
	return oldValue.IsTeacher, nil
}

// ClearIsTeacher clears the value of the "is_teacher" field.
func (m *UserMutation) ClearIsTeacher() {
	m.is_teacher = nil
	m.clearedFields[user.FieldIsTeacher] = struct{}{}
}

// IsTeacherCleared returns if the "is_teacher" field was cleared in this mutation.
func (m *UserMutation) IsTeacherCleared() bool {
	_, ok := m.clearedFields[user.FieldIsTeacher]
	return ok
}

// ResetIsTeacher resets all changes to the "is_teacher" field.
func (m *UserMutation) ResetIsTeacher() {
	m.is_teacher = nil
	delete(m.clearedFields, user.FieldIsTeacher)
}

// SetStars sets the "stars" field.
func (m *UserMutation) SetStars(i int32) {
	m.stars = &i
	m.addstars = nil
}

// Stars returns the value of the "stars" field in the mutation.
func (m *UserMutation) Stars() (r int32, exists bool) {
	v := m.stars
	if v == nil {
		return
	}
	return *v, true
}

// OldStars returns the old "stars" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStars(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStars is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStars requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStars: %w", err)
	}
	return oldValue.Stars, nil
}

// AddStars adds i to the "stars" field.
func (m *UserMutation) AddStars(i int32) {
	if m.addstars != nil {
		*m.addstars += i
	} else {
		m.addstars = &i
	}
}

// AddedStars returns the value that was added to the "stars" field in this mutation.
func (m *UserMutation) AddedStars() (r int32, exists bool) {
	v := m.addstars
	if v == nil {
		return
	}
	return *v, true
}

// ClearStars clears the value of the "stars" field.
func (m *UserMutation) ClearStars() {
	m.stars = nil
	m.addstars = nil
	m.clearedFields[user.FieldStars] = struct{}{}
}

// StarsCleared returns if the "stars" field was cleared in this mutation.
func (m *UserMutation) StarsCleared() bool {
	_, ok := m.clearedFields[user.FieldStars]
	return ok
}

// ResetStars resets all changes to the "stars" field.
func (m *UserMutation) ResetStars() {
	m.stars = nil
	m.addstars = nil
	delete(m.clearedFields, user.FieldStars)
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.real_name != nil {
		fields = append(fields, user.FieldRealName)
	}
	if m.mobile != nil {
		fields = append(fields, user.FieldMobile)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.is_teacher != nil {
		fields = append(fields, user.FieldIsTeacher)
	}
	if m.stars != nil {
		fields = append(fields, user.FieldStars)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldPassword:
		return m.Password()
	case user.FieldRealName:
		return m.RealName()
	case user.FieldMobile:
		return m.Mobile()
	case user.FieldEmail:
		return m.Email()
	case user.FieldIsTeacher:
		return m.IsTeacher()
	case user.FieldStars:
		return m.Stars()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldRealName:
		return m.OldRealName(ctx)
	case user.FieldMobile:
		return m.OldMobile(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldIsTeacher:
		return m.OldIsTeacher(ctx)
	case user.FieldStars:
		return m.OldStars(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldRealName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRealName(v)
		return nil
	case user.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldIsTeacher:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsTeacher(v)
		return nil
	case user.FieldStars:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStars(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addstars != nil {
		fields = append(fields, user.FieldStars)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldStars:
		return m.AddedStars()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldStars:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStars(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldRealName) {
		fields = append(fields, user.FieldRealName)
	}
	if m.FieldCleared(user.FieldMobile) {
		fields = append(fields, user.FieldMobile)
	}
	if m.FieldCleared(user.FieldEmail) {
		fields = append(fields, user.FieldEmail)
	}
	if m.FieldCleared(user.FieldIsTeacher) {
		fields = append(fields, user.FieldIsTeacher)
	}
	if m.FieldCleared(user.FieldStars) {
		fields = append(fields, user.FieldStars)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldRealName:
		m.ClearRealName()
		return nil
	case user.FieldMobile:
		m.ClearMobile()
		return nil
	case user.FieldEmail:
		m.ClearEmail()
		return nil
	case user.FieldIsTeacher:
		m.ClearIsTeacher()
		return nil
	case user.FieldStars:
		m.ClearStars()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldRealName:
		m.ResetRealName()
		return nil
	case user.FieldMobile:
		m.ResetMobile()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldIsTeacher:
		m.ResetIsTeacher()
		return nil
	case user.FieldStars:
		m.ResetStars()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
