package database

import (
	"slices"
	"fmt"
	"strings"
)

func (t *Table) validate() error {
	if strings.TrimSpace(t.name) == "" {
		return fmt.Errorf("table name cannot be empty")
	}
	
	if !validIdentifier(t.name) {
		return fmt.Errorf("invalid table name: %q", t.name)
	}

	if len(t.columns) == 0 {
		return fmt.Errorf("table %q must have at least one column", t.name)
	}
	
	if t.primaryCount > 1{
		return fmt.Errorf("table %q can not have more than one primary key", t.name)
	}
	

	for _, column := range t.columns {
		if !validIdentifier(column.name) {
			return fmt.Errorf(
				"invalid column name: %q",
				column.name,
			)
		}

		if column.columnType == "" {
			return fmt.Errorf(
				"column %q has no type",
				column.name,
			)
		}
		if column.composedPrimaryKey{

			if len(column.composedPrimary) < 2{
				return fmt.Errorf(
					"composedPrimary %q takes more than one column",
					column.name,
				)
			}
		}
		if column.autoInc {
			switch column.columnType {
			case "INT", "BIGINT", "SMALLINT":
				if !isIntegerType(column.columnType){
					return fmt.Errorf(
						"unsupported column type %q for column %q",
						column.columnType,
						column.name,
					)
				}
			default:
				return fmt.Errorf(
					"column %q cannot be AUTO_INCREMENT",
					column.name,
				)
			}
		}
		if column.primary && column.nullable {
			return fmt.Errorf(
				"primary key column %q cannot be nullable",
				column.name,
			)
		}

		if err := validateColumnType(column); err != nil {
			return err
		}
		
	}

	return t.validateColumns()
}
func (t *Table) validateEmptyIndex(name string, columns ...string) error{
	if name == ""{
		return fmt.Errorf("name of an index can't be empty at %q table", t.name)
	}
	if len(columns) < 1{
		return fmt.Errorf("columns can't be empty at %q table", t.name)
	}
	if slices.Contains(columns, "") {
		return fmt.Errorf("name of an column can't be empty at %q table", t.name)
	}
	return nil
}
