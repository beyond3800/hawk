package database

import "fmt"

type DropForeignKeyBuilder struct {
	table *Table
	name  string
}

func (t *Table) DropForeignKey(name string) *DropForeignKeyBuilder {
	return &DropForeignKeyBuilder{
		table: t,
		name:  name,
	}
}

func (d *DropForeignKeyBuilder) toSQL() string {
	return fmt.Sprintf(
		"ALTER TABLE %s DROP FOREIGN KEY %s",
		d.table.name,
		d.name,
	)
}

func (d *DropForeignKeyBuilder) Execute() error {
	_, err := d.table.db.Exec(d.toSQL())

	return err
}

// usage

// type RemoveUserForeignKey struct{}

// func (m RemoveUserForeignKey) Up() error {
// 	table := schema.Create("posts")

// 	return table.
// 		DropForeignKey("posts_user_id_foreign").
// 		Execute()
// }

// func (m RemoveUserForeignKey) Down() error {
// 	table := schema.Create("posts")

// 	return table.
// 		ForeignID("user_id").
// 		Constrained("users").
// 		Execute()
// }