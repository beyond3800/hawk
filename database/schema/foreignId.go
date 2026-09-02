package database

type foreignKey struct {
	Name          string
	Columns       []string
	References    string
	Table         string
	OnDelete      string
	OnUpdate      string
}

type ForeignIDBuilder struct {
    *ColumnBuilder
	foreignKey *foreignKey
}

type ActionBuilder struct {
	foreignIDBuilder       *ForeignIDBuilder
	actiontype             string
}

func (t *Table) ForeignID(name string) *ForeignIDBuilder {
    column := Column{
        name:       name,
        columnType: "BIGINT",
        nullable:   true,
    }

    t.columns = append(t.columns, column)

    return &ForeignIDBuilder{
        ColumnBuilder: &ColumnBuilder{
            table:  t,
            column: &t.columns[len(t.columns)-1],
        },
    }
}

func (f *ForeignIDBuilder) Constrained(table string) *ForeignIDBuilder {
    fk := foreignKey{
        Columns:    []string{f.column.name},
        Table:      table,
        References: "id",
    }

    f.table.foreignKeys = append(f.table.foreignKeys, fk)
	f.foreignKey = &f.table.foreignKeys[len(f.table.foreignKeys)-1]
    return f
}

func (f *ForeignIDBuilder) Name(indexName string) *ForeignIDBuilder{

	f.foreignKey.Name = indexName
	return f
}
func (f *ForeignIDBuilder) References(id string) *ForeignIDBuilder {
	f.foreignKey.References = id
    return f
}

func (f *ForeignIDBuilder) OnDelete() *ActionBuilder {
	if f.foreignKey != nil {
		f.foreignKey.OnDelete = "NO ACTION"
	}

	return &ActionBuilder{
		foreignIDBuilder: f,
		actiontype: "DELETE",
	}
}

func (f *ForeignIDBuilder) OnUpdate() *ActionBuilder {
	if f.foreignKey != nil {
		f.foreignKey.OnUpdate = "NO ACTION"
	}

	return &ActionBuilder{
		foreignIDBuilder: f,
		actiontype: "UPDATE",
	}
}

func (a *ActionBuilder) Cascade() *ForeignIDBuilder{

	switch a.actiontype {
	case "UPDATE":
		a.foreignIDBuilder.foreignKey.OnUpdate = "CASCADE"
	case "DELETE":
		a.foreignIDBuilder.foreignKey.OnDelete = "CASCADE"
	}
	return a.foreignIDBuilder
}
func (a *ActionBuilder) Restrict() *ForeignIDBuilder{

	switch a.actiontype {
	case "UPDATE":
		a.foreignIDBuilder.foreignKey.OnUpdate = "RESTRICT"
	case "DELETE":
		a.foreignIDBuilder.foreignKey.OnDelete = "RESTRICT"
	}
	return a.foreignIDBuilder
}
func (a *ActionBuilder) NoAction() *ForeignIDBuilder{

	switch a.actiontype {
	case "UPDATE":
		a.foreignIDBuilder.foreignKey.OnUpdate = "NO ACTION"
	case "DELETE":
		a.foreignIDBuilder.foreignKey.OnDelete = "NO ACTION"
	}
	return a.foreignIDBuilder
}
func (a *ActionBuilder) SetNull() *ForeignIDBuilder{

	switch a.actiontype {
	case "UPDATE":
		a.foreignIDBuilder.foreignKey.OnUpdate = "SET NULL"
	case "DELETE":
		a.foreignIDBuilder.foreignKey.OnDelete = "SET NULL"
	}
	return a.foreignIDBuilder
}