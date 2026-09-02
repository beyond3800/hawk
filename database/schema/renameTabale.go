package database

import "fmt"

type RenameTableBuilder struct {
	table *Table
	name  string
}

func (t *Table) Rename(name string) *RenameTableBuilder {
	return &RenameTableBuilder{
		table: t,
		name:  name,
	}
}

func (r *RenameTableBuilder) toSQL() string {
	return fmt.Sprintf(
		"ALTER TABLE %s RENAME TO %s",
		r.table.name,
		r.name,
	)
}

func (r *RenameTableBuilder) Execute() error {
	_, err := r.table.db.Exec(r.toSQL())

	return err
}

// example
// type RenameUsersTable struct{}

// func (m RenameUsersTable) Up() error {
// 	table := schema.Create("users")

// 	return table.Rename("customers").Execute()
// }

// func (m RenameUsersTable) Down() error {
// 	table := schema.Create("customers")

// 	return table.Rename("users").Execute()
// }