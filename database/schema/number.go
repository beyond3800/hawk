package database


func (t *Table) Decimal(name string, precision int, scale int) *ColumnBuilder {

	column := Column{
		name:       name,
		columnType: "DECIMAL",
		typeArgs:   []any{
			precision,
			scale,
		},
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}

func (t *Table) Integer(name string) *ColumnBuilder {
	column := Column{
		name:       name,
		columnType: "INT",
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}