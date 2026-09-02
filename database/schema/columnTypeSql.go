package database

import (
	"fmt"
	"strings"
)

func columnTypeSQL(column Column) (string, error) {
	switch column.columnType {

	case "VARCHAR":
		length := column.typeArgs[0].(int)

		return fmt.Sprintf(
			"VARCHAR(%d)",
			length,
		), nil

	case "DECIMAL":
		precision := column.typeArgs[0].(int)
		scale := column.typeArgs[1].(int)

		return fmt.Sprintf(
			"DECIMAL(%d,%d)",
			precision,
			scale,
		), nil
	case "ENUM":
		var enumValues []string
		seen := make(map[string]bool)

		if len(column.typeArgs) == 0 {
			return "",fmt.Errorf("enum column %q must have at least one value", column.name)
		}
		for _, value := range column.typeArgs {
			stringValue := value.(string)
			if stringValue == ""{
				return  "", fmt.Errorf("enum column can't be empty %q", column.name)
			}
			if seen[stringValue] {
			
				return "",
				 fmt.Errorf(
					"duplicate enum value %q for column %q",
					value,
					column.name,
				)
			}
			seen[stringValue] = true
			enumValues = append(enumValues, "'"+stringValue+"'")
		}
		return fmt.Sprintf(
			"ENUM(%s)",
			strings.Join(enumValues, ", "),
		), nil
	case "SET":
		var setValues []string
		seen := make(map[string]bool)

		if len(column.typeArgs) == 0 {
			return "",fmt.Errorf("set column %q must have at least one value", column.name)
		}
		for _, value := range column.typeArgs {
			stringValue := value.(string)
			if stringValue == ""{
				return  "", fmt.Errorf("set column can't be empty %q", column.name)
			}
			if seen[stringValue] {
			
				return "",
				 fmt.Errorf(
					"duplicate set value %q for column %q",
					value,
					column.name,
				)
			}
			seen[stringValue] = true
			setValues = append(setValues, "'"+stringValue+"'")
		}
		return fmt.Sprintf(
			"SET(%s)",
			strings.Join(setValues, ", "),
		), nil
	default:
		return column.columnType, nil
	}
}