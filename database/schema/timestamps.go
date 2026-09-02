package database


func (t *Table) Timestamps() {

	t.columns = append(t.columns, Column{
		name:       "created_at",
		columnType: "TIMESTAMP",
		defaultExpr:  "CURRENT_TIMESTAMP",
	})
	t.columns = append(t.columns, Column{
		name:       "updated_at",
		columnType: "TIMESTAMP",
		defaultExpr:  "CURRENT_TIMESTAMP",
		onUpdate:     "CURRENT_TIMESTAMP",
	})
}