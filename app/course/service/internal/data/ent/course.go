// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"microservice/app/course/service/internal/data/ent/course"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Course is the model entity for the Course schema.
type Course struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// TeacherID holds the value of the "teacher_id" field.
	TeacherID int64 `json:"teacher_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Course) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case course.FieldID, course.FieldTeacherID:
			values[i] = new(sql.NullInt64)
		case course.FieldTitle, course.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Course", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Course fields.
func (c *Course) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case course.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case course.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case course.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case course.FieldTeacherID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field teacher_id", values[i])
			} else if value.Valid {
				c.TeacherID = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Course.
// Note that you need to call Course.Unwrap() before calling this method if this Course
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Course) Update() *CourseUpdateOne {
	return (&CourseClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Course entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Course) Unwrap() *Course {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Course is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Course) String() string {
	var builder strings.Builder
	builder.WriteString("Course(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", title=")
	builder.WriteString(c.Title)
	builder.WriteString(", description=")
	builder.WriteString(c.Description)
	builder.WriteString(", teacher_id=")
	builder.WriteString(fmt.Sprintf("%v", c.TeacherID))
	builder.WriteByte(')')
	return builder.String()
}

// Courses is a parsable slice of Course.
type Courses []*Course

func (c Courses) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
