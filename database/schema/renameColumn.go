package database

import "fmt"

type RenameColumnBuilder struct {
	table *Table
	from  string
	to    string
}

func (t *Table) RenameColumn(from string, to string) *RenameColumnBuilder {
	return &RenameColumnBuilder{
		table: t,
		from:  from,
		to:    to,
	}
}

func (r *RenameColumnBuilder) toSQL() string {
	return fmt.Sprintf(
		"ALTER TABLE %s RENAME COLUMN %s TO %s",
		r.table.name,
		r.from,
		r.to,
	)
}

func (r *RenameColumnBuilder) Execute() error {
	_, err := r.table.db.Exec(r.toSQL())

	return err
}

// doc

// type RenameUserName struct{}

// func (m RenameUserName) Up() error {
// 	table := schema.Create("users")

// 	return table.RenameColumn("name", "username").Execute()
// }

// func (m RenameUserName) Down() error {
// 	table := schema.Create("users")

// 	return table.RenameColumn("username", "name").Execute()
// }