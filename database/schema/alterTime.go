package database


func (a *ColumnBuilder) Date() *ColumnBuilder {
	a.column.columnType = "DATE"
	a.column.typeArgs = nil

	return a
}

func (a *ColumnBuilder) DateTime() *ColumnBuilder {
	a.column.columnType = "DATETIME"
	a.column.typeArgs = nil

	return a
}

func (a *ColumnBuilder) Time() *ColumnBuilder {
	a.column.columnType = "TIME"
	a.column.typeArgs = nil

	return a
}

func (a *ColumnBuilder) Timestamp() *ColumnBuilder {
	a.column.columnType = "TIMESTAMP"
	a.column.typeArgs = nil

	return a
}

