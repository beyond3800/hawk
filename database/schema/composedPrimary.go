package database


func (t *Table) ComposedPrimary(columns ...string) *ColumnBuilder {
	t.primaryCount++

	column := &Column{
		composedPrimary: columns,
		composedPrimaryKey: true,
	}
	return &ColumnBuilder{
		table : t,
		column: column,
	}
}
func (t *Table) ID(){
	t.BigInt("id").Primary().AutoIncrement()
}