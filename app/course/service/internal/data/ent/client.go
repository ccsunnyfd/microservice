// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"microservice/app/course/service/internal/data/ent/migrate"

	"microservice/app/course/service/internal/data/ent/course"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Course is the client for interacting with the Course builders.
	Course *CourseClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Course = NewCourseClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Course: NewCourseClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Course: NewCourseClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Course.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Course.Use(hooks...)
}

// CourseClient is a client for the Course schema.
type CourseClient struct {
	config
}

// NewCourseClient returns a client for the Course from the given config.
func NewCourseClient(c config) *CourseClient {
	return &CourseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `course.Hooks(f(g(h())))`.
func (c *CourseClient) Use(hooks ...Hook) {
	c.hooks.Course = append(c.hooks.Course, hooks...)
}

// Create returns a create builder for Course.
func (c *CourseClient) Create() *CourseCreate {
	mutation := newCourseMutation(c.config, OpCreate)
	return &CourseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Course entities.
func (c *CourseClient) CreateBulk(builders ...*CourseCreate) *CourseCreateBulk {
	return &CourseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Course.
func (c *CourseClient) Update() *CourseUpdate {
	mutation := newCourseMutation(c.config, OpUpdate)
	return &CourseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CourseClient) UpdateOne(co *Course) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourse(co))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CourseClient) UpdateOneID(id int64) *CourseUpdateOne {
	mutation := newCourseMutation(c.config, OpUpdateOne, withCourseID(id))
	return &CourseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Course.
func (c *CourseClient) Delete() *CourseDelete {
	mutation := newCourseMutation(c.config, OpDelete)
	return &CourseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CourseClient) DeleteOne(co *Course) *CourseDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CourseClient) DeleteOneID(id int64) *CourseDeleteOne {
	builder := c.Delete().Where(course.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CourseDeleteOne{builder}
}

// Query returns a query builder for Course.
func (c *CourseClient) Query() *CourseQuery {
	return &CourseQuery{
		config: c.config,
	}
}

// Get returns a Course entity by its id.
func (c *CourseClient) Get(ctx context.Context, id int64) (*Course, error) {
	return c.Query().Where(course.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CourseClient) GetX(ctx context.Context, id int64) *Course {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CourseClient) Hooks() []Hook {
	return c.hooks.Course
}
