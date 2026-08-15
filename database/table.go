package database


func (db *DB) Table(name string) *Builder {
	return &Builder{
		db: db,
		table: name,
	}
}
func (b *Builder) Table(name string) *Builder {
	return &Builder{
		db:    b.db,
		tx:    b.tx,
		table: name,
	}
}