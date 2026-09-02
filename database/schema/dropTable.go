package database

import "fmt"

type DropTableBuilder struct {
	table *Table
}

func (t *Table) Drop() *DropTableBuilder {
	return &DropTableBuilder{
		table: t,
	}
}

func (d *DropTableBuilder) toSQL() string {
	return fmt.Sprintf(
		"DROP TABLE %s",
		d.table.name,
	)
}

func (d *DropTableBuilder) Execute() error {
	_, err := d.table.db.Exec(d.toSQL())

	return err
}

// example

// type DropUsersTable struct{}

// func (m DropUsersTable) Up() error {
// 	table := schema.Create("users")

// 	return table.Drop().Execute()
// }

// func (m DropUsersTable) Down() error {
// 	table := schema.Create("users")

// 	table.String("name", 255)

// 	return table.Create()
// }