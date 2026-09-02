package database

func (b *Builder) WithTrashed() *Builder {
	hasColumn, err := b.db.HasColumn(b.table, "deleted_at")

	if err != nil || !hasColumn {
		return b
	}

	b.withTrash = true
	b.onlyTrash = false

	return b
}

func (b *Builder) OnlyTrashed() *Builder {
	hasColumn, err := b.db.HasColumn(b.table, "deleted_at")

	if err != nil || !hasColumn {
		return b
	}

	b.onlyTrash = true
	b.withTrash = false

	return b
}