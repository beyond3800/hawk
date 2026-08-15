package validation

import (
	"reflect"
	"testing"
)

func TestRequired(t *testing.T) {
    v := &Validator{
        Errors: make(Errors),
    }

    field := reflect.StructField{
        Name: "Name",
    }

    v.Required(field, "")

    if _, ok := v.Errors["Name"]; !ok {
        t.Fatal("expected validation error")
    }
}

func TestMinLength(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	field := map[string]any{
		"Name": "abc",
	}

	v.Min("Name", field["Name"], "5")

	err, ok := v.Errors["Name"]

	if !ok {
		t.Fatal("expected validation error")
	}

	if err == "" {
		t.Fatal("expected validation error message")
	}
}

func TestMinLengthPasses(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	field := map[string]any{
		"Name": "abcdef",
	}

	v.Min("Name", field["Name"], "5")

	if _, ok := v.Errors["Name"]; ok {
		t.Fatalf("expected no validation error, got %s", v.Errors["Name"])
	}
}

func TestMaxLength(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	field := map[string]any{
		"Name": "abcdef",
	}

	v.Max("Name", field["Name"], "5")

	if _, ok := v.Errors["Name"]; !ok {
		t.Fatal("expected validation error")
	}
}

func TestMaxLengthPasses(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	field := map[string]any{
		"Name": "abc",
	}

	v.Max("Name", field["Name"], "5")

	if _, ok := v.Errors["Name"]; ok {
		t.Fatalf("expected no validation error, got %s", v.Errors["Name"])
	}
}

func TestEmail(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	field := reflect.StructField{
		Name: "Email",
	}

	v.Email(field.Name, "invalid-email")

	if _, ok := v.Errors["Email"]; !ok {
		t.Fatal("expected validation error")
	}
}

func TestEmailPasses(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	field := reflect.StructField{
		Name: "Email",
	}

	v.Email(field.Name, "adam@test.com")

	if _, ok := v.Errors["Email"]; ok {
		t.Fatalf("expected no validation error, got %s", v.Errors["Email"])
	}
}

func TestMultipleValidationErrors(t *testing.T) {
	v := &Validator{
		Errors: make(Errors),
	}

	nameField := reflect.StructField{
		Name: "Name",
	}

	emailField := reflect.StructField{
		Name: "Email",
	}

	v.Required(nameField, "")
	v.Email(emailField.Name, "invalid-email")

	if _, ok := v.Errors["Name"]; !ok {
		t.Fatal("expected Name validation error")
	}

	if _, ok := v.Errors["Email"]; !ok {
		t.Fatal("expected Email validation error")
	}
}