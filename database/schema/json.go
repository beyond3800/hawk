package database


func (t *Table) JSON(name string) *ColumnBuilder {
	column := Column{
		name:       name,
		columnType: "JSON",
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}