package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Table struct {
	name           string
	db             *sql.DB
	columns        []Column
	primaryCount   int

	indexes        []index
	uniques        []index
	fullTexts      []index
	dropIndexes    []string
	checks         []check
	foreignKeys    []foreignKey
	
	err            []error
}

func (t *Table) columnType(name string) string {
    for _, column := range t.columns {
        if column.name == name {
            return column.columnType
        }
    }

    return ""
}

func (t *Table) indexNameExists(name string) bool {
    for _, index := range t.indexes {
        if index.Name != "" && index.Name == name {
            return true
        }
    }

    for _, index := range t.uniques {
        if index.Name != "" && index.Name == name {
            return true
        }
    }

    for _, index := range t.fullTexts {
        if index.Name != "" && index.Name == name {
            return true
        }
    }

    return false
}

func validateIndexColumns(columns []string) error {
    seen := make(map[string]bool)

    for _, column := range columns {
        if seen[column] {
            return fmt.Errorf(
                "column %q appears more than once in index",
                column,
            )
        }

        seen[column] = true
    }

    return nil
}

func (t *Table) validateColumns() error {
	seen := make(map[string]bool)

	for _, column := range t.columns {
		name := strings.TrimSpace(column.name)

		if !isSupportedColumnType(column.columnType) {
			fmt.Println(column.name,column.columnType)
			return fmt.Errorf(
				"unsupported column type %q for column %q",
				column.columnType,
				name,
			)
		}
		
		if name == "" {
			return fmt.Errorf("column name cannot be empty")
		}

		if seen[name] {
			return fmt.Errorf(
				"duplicate column %q in table %q",
				name,
				t.name,
			)
		}

		seen[name] = true
	}

	if len(t.indexes) > 0{
		for _, index := range t.indexes{
			for _, column := range index.Columns{
				if !seen[column]{
					return fmt.Errorf("This column does not exist in table %s", t.name)
				}
			}
			if err := validateIndexColumns(index.Columns); err != nil {
				return err
			}
			 
		}
	}
	if len(t.uniques) > 0{
		for _, index := range t.uniques{
			for _, column := range index.Columns{
				if !seen[column]{
					return fmt.Errorf("This column does not exist in table %s", t.name)
				}
			}
			if err := validateIndexColumns(index.Columns); err != nil {
				return err
			}
		}
	}
	if len(t.fullTexts) > 0{
		for _, index := range t.fullTexts{
			for _, column := range index.Columns{
				if !seen[column]{
					return fmt.Errorf("This column does not exist in table %s", t.name)
				}
				columnType := t.columnType(column)
				switch columnType {
				case "CHAR",
					"VARCHAR",
					"TINYTEXT",
					"TEXT",
					"MEDIUMTEXT",
					"LONGTEXT":
				default:
					return fmt.Errorf(
						"column %q of type %q can not have full text index",
						column,
						columnType,
					)
				}
				if err := validateIndexColumns(index.Columns); err != nil {
					return err
				}
				
				// for _,column := range t.columns{
				// 	if t.indexNameExists(column.fullTextName){
				// 		return fmt.Errorf(
				// 			"index %q already exists on table %q",
				// 			column.fullTextName,
				// 			t.name,
				// 		)
				// 	}
				// }
			}
		}

	}
	if len(t.foreignKeys) > 0{
		for _, foreignKey := range t.foreignKeys {
			for _, column := range foreignKey.Columns {
				if !seen[column] {
					return fmt.Errorf(
						"foreign key column %q does not exist in table %q",
						column,
						t.name,
					)
				}
			}
		}
	}
	
	return nil
}

func (t *Table) Truncate() error {
	_, err := t.db.Exec(
		fmt.Sprintf("TRUNCATE TABLE %s", t.name),
	)

	return err
}



