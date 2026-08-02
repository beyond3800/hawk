package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/beyond3800/hawk/internal/bootstrap/database"
)

type Errors map[string]string
type Validator struct {
    Errors Errors
}

func (v *Validator) Required(field reflect.StructField, value any){

    rv := reflect.ValueOf(value)
        switch rv.Kind() {

    case reflect.String:
        if rv.String() == "" {
            v.Errors[field.Name] = field.Name + " is required"
        }

    case reflect.Slice, reflect.Map:
        if rv.IsNil() || rv.Len() == 0 {
                v.Errors[field.Name] = field.Name + " is required"
            }
    case reflect.Array:
        if rv.Len() == 0 {
            v.Errors[field.Name] = field.Name + " is required"
        }

    case reflect.Ptr, reflect.Interface:
        if rv.IsNil() {
            v.Errors[field.Name] = field.Name + " is required"
        }
    }
    
}
func (v *Validator) Email(field string, value any){
    str, ok := value.(string)
    if !ok {
        v.Errors[field] = field + " must be a string"
        return
    }
    if !emailRegex.MatchString(str) {
        v.Errors[field] = field + " must be a valid email"
    }
}
func (v *Validator) Min(field string, value any, minValue string) {
	minInt, err := strconv.Atoi(minValue)
	if err != nil {
		v.Errors[field] = field + " invalid min value"
		return
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {

	case reflect.String:
		if len([]rune(rv.String())) < minInt {
			v.Errors[field] = field + " can't be less than " + minValue
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv.Int() < int64(minInt) {
			v.Errors[field] = field + " can't be less than " + minValue
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if rv.Uint() < uint64(minInt) {
			v.Errors[field] = field + " can't be less than " + minValue
		}

	case reflect.Float32, reflect.Float64:
		if rv.Float() < float64(minInt) {
			v.Errors[field] = field + " can't be less than " + minValue
		}

	default:
		v.Errors[field] = field + " does not support min validation"
	}
}
func (v *Validator) Max(field string, value any, maxValue string) {
	maxInt, err := strconv.Atoi(maxValue)
	if err != nil {
		v.Errors[field] = field + " invalid max value"
		return
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {

	case reflect.String:
		if len([]rune(rv.String())) > maxInt {
			v.Errors[field] = field + " can't be greater than " + maxValue
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rv.Int() > int64(maxInt) {
			v.Errors[field] = field + " can't be greater than " + maxValue
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if rv.Uint() > uint64(maxInt) {
			v.Errors[field] = field + " can't be greater than " + maxValue
		}

	case reflect.Float32, reflect.Float64:
		if rv.Float() > float64(maxInt) {
			v.Errors[field] = field + " can't be greater than " + maxValue
		}

	default:
		v.Errors[field] = field + " does not support max validation"
	}
}
func (v *Validator) Unique(field string, value string, table string){
    var count int
    ok := stringRegex.MatchString(table)
    if !ok {
        v.Errors[field] = field + " is not a valid table name"
        return
    }
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, field)
    err := database.DB.QueryRow(
        query,
        value,
    ).Scan(&count)
    if err != nil {
        v.Errors[field] = "database validation failed"
    }
    if count > 0 {
        v.Errors[field] = field + " is already taken"
    }
}

func Validate(data any) (Errors, error) {
    v := &Validator{
        Errors: make(map[string]string),
    }
    val := reflect.ValueOf(data)
    if val.Kind() == reflect.Ptr{
        val = val.Elem()
    }
    if val.Kind() != reflect.Struct{
        v.Errors["data"] = "validation requires a struct"
        return v.Errors, fmt.Errorf("validation requires a struct")
    }
    structType := val.Type()
    for i := 0; i < structType.NumField(); i++ {

        field := structType.Field(i)
        // fieldValue := val.Field(i)
        validateRule := field.Tag.Get("validate")
        if validateRule != "" {
            // fmt.Println(fieldValue)
            rules := strings.Split(validateRule, "|")
            for _,rule := range rules{
                switch rule {
                case "required":
                    v.Required(field, val.Field(i).Interface())
                case "email":
                    v.Email(field.Name,val.Field(i).Interface())
                case "unique":
                    uniqueRule := strings.Split(field.Tag.Get("unique"),",")
                    if len(uniqueRule) != 2 {
                        v.Errors[field.Name] = field.Name + " unique rule is invalid"
                        continue
                    }
                    v.Unique(uniqueRule[1],val.Field(i).Interface().(string),uniqueRule[0])
                }
                if strings.HasPrefix(rule, "min:"){
                    minValue := strings.TrimPrefix(rule, "min:")
                    v.Min(field.Name,val.Field(i).Interface(),minValue)
                }
                if strings.HasPrefix(rule, "max:"){
                    maxValue := strings.TrimPrefix(rule, "max:")
                    v.Max(field.Name,val.Field(i).Interface(),maxValue)
                }
                
            } 
        }

    }
    if len(v.Errors) > 0{
        return v.Errors, fmt.Errorf("Validation Error")
    }
    return nil, nil
}