package database


func (db *DB) Table(name string) *Builder {
	hasSoftDeletes, err := db.hasSoftDeletes(name)
	return &Builder{
		db: db,
		table: name,
		hasSoftDeletes: hasSoftDeletes,
		schemaErr: err,
	}
}
func (b *Builder) Table(name string) *Builder {
	hasSoftDeletes, err := b.db.hasSoftDeletes(name)
	return &Builder{
		db:    b.db,
		tx:    b.tx,
		table: name,
		hasSoftDeletes: hasSoftDeletes,
		schemaErr: err,
	}
}