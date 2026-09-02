package database


func (t *Table) SoftDeletes() {
	
	t.columns = append(t.columns, Column{
		name:       "deleted_at",
		columnType: "TIMESTAMP",
		nullable:   true,
	})
}
