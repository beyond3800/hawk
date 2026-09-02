package database

import "strings"

func (b *Builder) Where(column string, value any) *Builder {

	b.wheres = append(
		b.wheres,
		column+" = ?",
	)

	b.bindings = append(
		b.bindings,
		value,
	)

	return b
}

func (b *Builder) OrWhere(column string, value any) *Builder {

	condition := column + " = ?"

	if len(b.wheres) == 0 {
		b.wheres = append(b.wheres, condition)
	} else {
		b.wheres = append(b.wheres, "OR "+condition)
	}

	b.bindings = append(b.bindings, value)

	return b
}

func (b *Builder) buildWhere() string {
	var conditions []string

	if b.hasSoftDeletes && !b.withTrash {
		if b.onlyTrash {
			conditions = append(conditions, "deleted_at IS NOT NULL")
		} else {
			conditions = append(conditions, "deleted_at IS NULL")
		}
	}

	conditions = append(conditions, b.wheres...)

	if len(conditions) == 0 {
		return ""
	}

	return " WHERE " + strings.Join(conditions, " AND ")
}