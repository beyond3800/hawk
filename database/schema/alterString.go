package database


func (c *ColumnBuilder) String(length int) *ColumnBuilder {
	c.column.columnType = "VARCHAR"
	c.column.typeArgs = []any{length}

	return c
}

func (c *ColumnBuilder) Char(length int) *ColumnBuilder {
	c.column.columnType = "CHAR"
	c.column.typeArgs = []any{length}

	return c
}

func (c *ColumnBuilder) Text() *ColumnBuilder {
	c.column.columnType = "TEXT"
	c.column.typeArgs = nil

	return c
}

func (c *ColumnBuilder) MediumText() *ColumnBuilder {
	c.column.columnType = "MEDIUMTEXT"
	c.column.typeArgs = nil

	return c
}

func (c *ColumnBuilder) LongText() *ColumnBuilder {
	c.column.columnType = "LONGTEXT"
	c.column.typeArgs = nil

	return c
}

func (c *ColumnBuilder) Enum(values ...string) *ColumnBuilder {


	c.column.columnType = "ENUM"
	c.column.typeArgs = []any{values}

	return c
}

func (c *ColumnBuilder) Set(values ...string) *ColumnBuilder {

	typeArgs := make([]any, len(values))

	for i, value := range values {
		typeArgs[i] = value
	}
	c.column.columnType = "SET"
	c.column.typeArgs = typeArgs

	return c
}