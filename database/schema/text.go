package database


func (t *Table) Text(name string) *ColumnBuilder{
	column := Column{
		name:       name,
		columnType: "TEXT",
		nullable: true,
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}
func (t *Table) Enum(name string, values ...string) *ColumnBuilder {

	typeArgs := make([]any, len(values))

	for i, value := range values {
		typeArgs[i] = value
	}

	column := Column{
		name:       name,
		columnType: "ENUM",
		typeArgs:   typeArgs,
		nullable:   true,
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}
func (t *Table) Set(name string, values ...string) *ColumnBuilder {

	typeArgs := make([]any, len(values))

	for i, value := range values {
		typeArgs[i] = value
	}

	column := Column{
		name:       name,
		columnType: "SET",
		typeArgs:   typeArgs,
		nullable:   true,
	}

	t.columns = append(t.columns, column)

	return &ColumnBuilder{
		table:  t,
		column: &t.columns[len(t.columns)-1],
	}
}