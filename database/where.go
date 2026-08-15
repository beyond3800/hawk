package database


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