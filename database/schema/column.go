package database

import (
	"fmt"
	"strings"
)


type Column struct {
	name                string
	columnType          string
	typeArgs            []any

	nullable            bool
	primary             bool
	autoInc             bool

	unique              bool
	uniqueName          string

	hasDefault          bool
	defaultValue        any

	defaultExpr         string
	onUpdate            string

	unsigned            bool
	comment             string

	composedPrimaryKey  bool
	composedPrimary     []string

	first               bool
	after               string
	operation           string
	renameTo            string

}

func formatDefault(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return "'" + strings.ReplaceAll(v, "'", "''") + "'", nil

	case bool:
		if v {
			return "TRUE", nil
		}
		return "FALSE", nil

	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return fmt.Sprintf("%v", v), nil

	case nil:
		return "NULL", nil

	default:
		return "", fmt.Errorf(
			"unsupported default value type %T",
			value,
		)
	}
}

type ColumnBuilder struct {
	table  *Table
	column *Column
}

func (c *ColumnBuilder) Nullable() *ColumnBuilder {
	c.column.nullable = true
	return c
}

func (c *ColumnBuilder) Required() *ColumnBuilder {
	c.column.nullable = false
	return c
}

func (c *ColumnBuilder) Unique(name ...string) *ColumnBuilder {
	c.column.unique = true
	return c
}

func (c *ColumnBuilder) Default(value any) *ColumnBuilder {
	c.column.hasDefault = true
	c.column.defaultValue = value
	return c
}

func (c *ColumnBuilder) DefaultExpr(expr string) *ColumnBuilder {
	c.column.defaultExpr = expr
	return c
}

func (c *ColumnBuilder) OnUpdate(expr string) *ColumnBuilder {
	c.column.onUpdate = expr
	return c
}

func (c *ColumnBuilder) AutoIncrement() *ColumnBuilder {
	c.column.autoInc = true
	return c
}

func (c *ColumnBuilder) Primary() *ColumnBuilder {
	c.column.primary = true
	c.column.nullable = false
	c.table.primaryCount++
	return c
}

func (c *ColumnBuilder) Index(name string) *ColumnBuilder{
    indx := index{
		Name: name,
		Columns: []string{c.column.name},
		Type: "INDEX",
	}
	c.table.err = append(c.table.err, c.table.addToIndex(indx))
	return c
}

func (c *ColumnBuilder) FullText(name string) *ColumnBuilder{
	indx :=index{
		Name: name,
		Columns: []string{c.column.name},
		Type: "FULLTEXT",
	}
	c.table.err = append(c.table.err, c.table.addToIndex(indx))
	return  c
}

func (c *ColumnBuilder) Unsigned() *ColumnBuilder {
	c.column.unsigned = true

	return c
}
func newColumnBuilder(table *Table, name, operation string) *ColumnBuilder{
	return &ColumnBuilder{
		table: table,
		column: &Column{
			name: name,
			nullable: true,
			operation: operation,
		},
	}
}




