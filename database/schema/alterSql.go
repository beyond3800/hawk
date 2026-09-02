package database

import (
	"fmt"
	"strings"
)

func (c *ColumnBuilder) alterSQL() (string, error) {
	var query strings.Builder

	query.WriteString("ALTER TABLE ")
	query.WriteString(c.table.name)
	query.WriteString(" ")

	switch c.column.operation {
	case "ADD":
		query.WriteString("ADD COLUMN ")
	case "MODIFY":
		query.WriteString("MODIFTY COLUMN ")
	default:
		return "",fmt.Errorf("this isn't allowed")
	}
	if c.column.renameTo != "" {
		query.WriteString("RENAME COLUMN ")
		query.WriteString(c.column.name)
		query.WriteString(" TO ")
		query.WriteString(c.column.renameTo)

		return query.String(), nil
	}

	query.WriteString("MODIFY COLUMN ")
	query.WriteString(c.column.name)
	query.WriteString(" ")

	query.WriteString(c.column.columnType)

	if len(c.column.typeArgs) > 0 {
		// Handle type arguments.
		switch c.column.columnType {
		case "VARCHAR", "CHAR":
			query.WriteString(
				fmt.Sprintf("(%v)", c.column.typeArgs[0]),
			)

		case "DECIMAL":
			query.WriteString(
				fmt.Sprintf(
					"(%v, %v)",
					c.column.typeArgs[0],
					c.column.typeArgs[1],
				),
			)

		case "ENUM", "SET":
			values := c.column.typeArgs[0].([]string)

			var enumValues []string

			for _, value := range values {
				enumValues = append(
					enumValues,
					"'"+value+"'",
				)
			}

			query.WriteString(
				fmt.Sprintf(
					"(%s)",
					strings.Join(enumValues, ", "),
				),
			)
		}
	}

	if c.column.unsigned {
		query.WriteString(" UNSIGNED")
	}

	if c.column.nullable {
		query.WriteString(" NULL")
	} else {
		query.WriteString(" NOT NULL")
	}

	if c.column.hasDefault {
		query.WriteString(" DEFAULT ")
		sql, err := formatDefault(c.column.defaultValue)
		if err != nil{
			return "", err
		}
		query.WriteString(sql)
	}

	if c.column.autoInc {
		query.WriteString(" AUTO_INCREMENT")
	}

	if c.column.comment != "" {
		query.WriteString(" COMMENT ")
		query.WriteString(
			fmt.Sprintf("'%s'", c.column.comment),
		)
	}

	if c.column.first {
		query.WriteString(" FIRST")
	} else if c.column.after != "" {
		query.WriteString(" AFTER ")
		query.WriteString(c.column.after)
	}

	return query.String(), nil
}