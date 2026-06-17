package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/beyond3800/hawk/lib"
	_ "golang.org/x/text/cases"
)

func Validate(obj any, fields ...string) []lib.FieldError{
	var errors []lib.FieldError
	val := reflect.ValueOf(obj)

	for _, field := range fields {
		fieldValue := val.FieldByName(strings.Title(field))
		
		if !fieldValue.IsValid() {
			errors = append(errors, lib.FieldError{
				FieldName:field,
				Message: fmt.Sprintf("Field %s does not exist", field),
			})
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			if strings.TrimSpace(fieldValue.String()) == "" {
				errors = append(errors, lib.FieldError{
					FieldName:field,
					Message: fmt.Sprintf("Field %s is empty", field),
				})
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldValue.Type().Kind() != reflect.Int{
				errors = append(errors, lib.FieldError{
					FieldName:field,
					Message: fmt.Sprintf("Field %s is supposed to be a number", field),
				})
			}
		case reflect.Float32, reflect.Float64:
			
			if fieldValue.Float() == 0.0 {
				errors = append(errors, lib.FieldError{
					FieldName:field,
					Message: fmt.Sprintf("Field %s is empty", field),
				})
			}
		case reflect.Bool:
			if !fieldValue.Bool() {
				errors = append(errors, lib.FieldError{
					FieldName:field,
					Message: fmt.Sprintf("Field %s does not exist", field),
				})
			}
		case reflect.Slice, reflect.Map:
			if fieldValue.Len() == 0 {
				errors = append(errors, lib.FieldError{
					FieldName:field,
					Message: fmt.Sprintf("Field %s does not exist", field),
				})
			}
		case reflect.Struct:
			// fmt.Println(reflect.Zero(fieldValue.Type()).Interface())
		// Check if struct is zero value
			if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
				errors = append(errors, lib.FieldError{
					FieldName: field,
					Message:   fmt.Sprintf("Field %s is empty", field),
				})
				
			} else {
				// Recursively validate struct fields
				t := fieldValue.Type()
				for i := 0; i < t.NumField(); i++ {
					subField := t.Field(i)
					subValue := fieldValue.Field(i)
					// Example: check if string fields are empty
					switch subValue.Kind(){
						case reflect.String:
							if strings.TrimSpace(subValue.String()) == ""{
								errors = append(errors, lib.FieldError{
									FieldName: fmt.Sprintf("%s.%s", field, subField.Name),
									Message:   fmt.Sprintf("Field %s.%s is empty", field, subField.Name),
								})
							}
					
					}
					
				}
			}

		default:
			// Optional: skip other types
		}
	}
	return errors
}


