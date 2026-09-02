package database


func (t *Table) String(name string, length int) *ColumnBuilder {
	column := Column{
		name:       name,
		columnType: "VARCHAR",
		typeArgs:   []any{length},
		nullable: true,
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}