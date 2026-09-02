package database

import "fmt"

var supportedColumnTypes = map[string]bool{
	"INT":       true,
	"BIGINT":    true,
	"SMALLINT":  true,
	"VARCHAR":   true,
	"MEDIUMTEXT":true,
	"LONGTEXT":  true,
	"TINYTEXT":  true,
	"CHAR":      true,
	"TEXT":      true,
	"ENUM":      true,
	"SET":       true,
	"BOOLEAN":   true,
	"FLOAT":     true,
	"DECIMAL":   true,
	"JSON":      true,
	"DATE":      true,
	"TIME":      true,
	"TIMESTAMP": true,
}

func isSupportedColumnType(columnType string) bool {
	return supportedColumnTypes[columnType]
}

func isIntegerType(columnType string) bool {
	switch columnType {
	case "INT", "BIGINT", "SMALLINT":
		return true
	default:
		return false
	}
}

func validateColumnType(column Column) error {
	switch column.columnType {

	case "VARCHAR":
		if len(column.typeArgs) != 1 {
			return fmt.Errorf(
				"VARCHAR column %q requires a length",
				column.name,
			)
		}

		length, ok := column.typeArgs[0].(int)
		if !ok || length <= 0 {
			return fmt.Errorf(
				"VARCHAR column %q requires a positive length",
				column.name,
			)
		}
	case "DECIMAL":
		if len(column.typeArgs) != 2 {
			return fmt.Errorf(
				"DECIMAL column %q requires precision and scale",
				column.name,
			)
		}

		precision, ok1 := column.typeArgs[0].(int)
		scale, ok2 := column.typeArgs[1].(int)

		if !ok1 || !ok2 {
			return fmt.Errorf(
				"DECIMAL column %q requires integer precision and scale",
				column.name,
			)
		}

		if precision <= 0 {
			return fmt.Errorf(
				"DECIMAL column %q requires positive precision",
				column.name,
			)
		}

		if scale < 0 || scale > precision {
			return fmt.Errorf(
				"invalid scale for DECIMAL column %q",
				column.name,
			)
		}
	}

	return nil
}