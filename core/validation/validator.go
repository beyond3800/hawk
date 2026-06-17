package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)


type Validator struct {
    error map[string]string
}

func (v *Validator) Required(field string, value any){

    if value == "" || value == nil{
        v.error[field] = field + " is required"
    }
}
func (v *Validator) Email(field string, value string){
    if !strings.Contains(value, "@") {
        v.error[field] = field + " must be a valid email"
    }
}
func (v *Validator) Min(field string, value string, minValue string){
    minInt, err := strconv.Atoi(minValue)
    if err != nil {
        v.error[field] = field  + " Invalid min value"
    }
    if value == ""{
        v.error[field] = field  + " is required" 
        return
    }
    valueLength := len([]rune(value))
    if valueLength < minInt{
        v.error[field] = field  + " can't be less than " + minValue
    }

}
func (v *Validator) Max(field string, value string, maxValue string){
    minInt, err := strconv.Atoi(maxValue)
    if err != nil {
        v.error[field] = field  + " Invalid min value"
    }
    if value == ""{
        v.error[field] = field  + " is required" 
        return
    }
    valueLength := len([]rune(value))
    if valueLength > minInt{
        v.error[field] = field  + " can't be greater than " + maxValue
    }

}
func (v *Validator) Unique(field string, value string, table string){

}

func Validate(data any) map[string]string {
    fmt.Println("validate")
    v := &Validator{
        error: make(map[string]string),
    }
    val := reflect.ValueOf(data)
    if val.Kind() == reflect.Ptr{
        val = val.Elem()
    }
    if val.Kind() != reflect.Struct{
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
                    v.Required(field.Name, val.Field(i).Interface())
                case "email":
                    v.Email(field.Name,val.Field(i).Interface().(string))
                }
                if strings.HasPrefix(rule, "min:"){
                    minValue := strings.TrimPrefix(rule, "min:")
                    v.Min(field.Name,val.Field(i).Interface().(string),minValue)
                }
                if strings.HasPrefix(rule, "max:"){
                    maxValue := strings.TrimPrefix(rule, "max:")
                    v.Max(field.Name,val.Field(i).Interface().(string),maxValue)
                }
            } 
        }

    }
    return v.error
}