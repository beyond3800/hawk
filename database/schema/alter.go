package database

import "fmt"

type AlterColumnBuilder struct {
	table  *Table
	column *Column

	renameTo string
	first    bool
	after    string
}

func (t *Table) AlterColumn(name string) *ColumnBuilder {
	return newColumnBuilder(t, name, "MODIFY")
}

func (c *ColumnBuilder) JSON() *ColumnBuilder {
	c.column.columnType = "JSON"
	c.column.typeArgs = nil

	return c
}

func (c *ColumnBuilder) Boolean() *ColumnBuilder {
	c.column.columnType = "BOOLEAN"
	c.column.typeArgs = nil

	return c
}

func (c *ColumnBuilder) DropDefault() *ColumnBuilder {
	c.column.hasDefault = false
	c.column.defaultValue = nil

	return c
}

func (c *ColumnBuilder) Rename(name string) *ColumnBuilder {
	c.column.renameTo = name

	return c
}

func (c *ColumnBuilder) First() *ColumnBuilder {
	c.column.first = true
	c.column.after = ""

	return c
}

func (c *ColumnBuilder) After(column string) *ColumnBuilder {
	c.column.after = column
	c.column.first = false

	return c
}

func (c *ColumnBuilder) Execute() error {
	query, err := c.alterSQL()

	if err != nil {
		return err
	}
	fmt.Println(query)
	_, err = c.table.db.Exec(query)

	return err
}