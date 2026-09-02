package database

import "fmt"

type DropColumnBuilder struct {
	table *Table
	name  string
}

func (t *Table) DropColumn(name string) *DropColumnBuilder {
	return &DropColumnBuilder{
		table: t,
		name:  name,
	}
}

func (d *DropColumnBuilder) toSQL() string {
	return fmt.Sprintf(
		"ALTER TABLE %s DROP COLUMN %s",
		d.table.name,
		d.name,
	)
}
func (d *DropColumnBuilder) Execute() error {
	_, err := d.table.db.Exec(d.toSQL())

	return err
}


// doc

// type DropUsersAge struct{}

// func (m DropUsersAge) Up() error {
// 	table := schema.Create("users")

// 	return table.DropColumn("age").Execute()
// }

// func (m DropUsersAge) Down() error {
// 	table := schema.Create("users")

// 	return table.AddColumn("age").
// 		Int().
// 		Nullable().
// 		Execute()
// }