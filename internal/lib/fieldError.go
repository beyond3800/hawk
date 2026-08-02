package lib

type FieldError struct{
	FieldName string `json:"field_name"`
	Message string `json:"message"`
}