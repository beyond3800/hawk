package database

import (
	"fmt"
	"strings"
)

func (s *Schema) Table(name string) *Table {
	return &Table{
		name:    name,
		db:      s.Db,
		columns: []Column{},
	}
}

func (t *Table) Create() error {

	columns := make([]string, 0, len(t.columns))
	
	// at column level
	for _, column := range t.columns {
		columnType,err := columnTypeSQL(column)

		if err != nil {
			return err
		}
		
		
		definition := fmt.Sprintf(
			"%s %s",
			column.name,
			columnType,
		)

		if column.autoInc {
			definition += " AUTO_INCREMENT"
		}

		if column.primary {
			if len(column.composedPrimary) > 1 {
				column := strings.Join(column.composedPrimary, ",")
				definition += fmt.Sprintf(" PRIMARY KEY (%s)", column)
			}
			definition += " PRIMARY KEY"
		}

		if !column.nullable {
			definition += " NOT NULL"
		}

		if column.unique {
			definition += " UNIQUE"
		}

		if column.hasDefault {
			defaultValue, err := formatDefault(column.defaultValue)
			if err != nil {
				return err
			}

			definition += " DEFAULT " + defaultValue
		}
		if column.defaultExpr != "" {
			definition += " DEFAULT " + column.defaultExpr
		}

		if column.onUpdate != "" {
			definition += " ON UPDATE " + column.onUpdate
		}
		columns = append(columns, definition)
	}
	if err := t.validate(); err != nil {
		return err
	}
	if len(t.err) > 0{
		for _,err := range t.err{
			if err != nil{
				return err
			}
		}
	}
	createdIndex, err :=  t.createIndexesQuery()
	if err != nil {
		return err
	}
	var query string
	if len(createdIndex)>0{
		query = fmt.Sprintf(
			"CREATE TABLE %s (%s,%s)",
			t.name,
			strings.Join(columns, ", "),
			strings.Join(createdIndex, ","),
		)
	}else{
		query = fmt.Sprintf(
			"CREATE TABLE %s (%s)",
			t.name,
			strings.Join(columns, ", "),
		)
	}

	_, err = t.db.Exec(query)
	if err != nil {
		return fmt.Errorf(
			"failed to create table %q: %w",
			t.name,
			err,
		)
	}

	return nil
}


