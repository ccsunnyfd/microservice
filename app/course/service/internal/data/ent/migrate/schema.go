// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "teacher_id", Type: field.TypeInt64, Nullable: true},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoursesTable,
	}
)

func init() {
}
