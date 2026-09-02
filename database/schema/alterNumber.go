package database

func (a *AlterColumnBuilder) TinyInt() *AlterColumnBuilder {
	a.column.columnType = "TINYINT"
	a.column.typeArgs = nil

	return a
}

func (a *AlterColumnBuilder) SmallInt() *AlterColumnBuilder {
	a.column.columnType = "SMALLINT"
	a.column.typeArgs = nil

	return a
}

func (a *AlterColumnBuilder) Int() *AlterColumnBuilder {
	a.column.columnType = "INT"
	a.column.typeArgs = nil

	return a
}

func (a *AlterColumnBuilder) BigInt() *AlterColumnBuilder {
	a.column.columnType = "BIGINT"
	a.column.typeArgs = nil

	return a
}
func (a *AlterColumnBuilder) Float() *AlterColumnBuilder {
	a.column.columnType = "FLOAT"
	a.column.typeArgs = nil

	return a
}

func (a *AlterColumnBuilder) Double() *AlterColumnBuilder {
	a.column.columnType = "DOUBLE"
	a.column.typeArgs = nil

	return a
}

func (a *AlterColumnBuilder) Decimal(precision, scale int) *AlterColumnBuilder {
	a.column.columnType = "DECIMAL"
	a.column.typeArgs = []any{precision, scale}

	return a
}