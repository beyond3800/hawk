package database

func (t *Table) BigInt(name string) *ColumnBuilder {
	column := Column{
		name:       name,
		columnType: "BIGINT",
		nullable: true,
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}

