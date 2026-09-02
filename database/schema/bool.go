package database



func (t *Table) Boolean(name string) *ColumnBuilder {
	column := Column{
		name:       name,
		columnType: "BOOLEAN",
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}

