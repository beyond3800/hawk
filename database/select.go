package database


func (b *Builder) Select(columns ...string) *Builder {
	b.columns = columns
	return b
}