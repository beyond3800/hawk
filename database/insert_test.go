package database

import (
	"testing"
)

func TestInsert(t *testing.T) {
	setupTestDB(t)
	type User struct {
		Name  string
		Email string
	}
	var user User

	user.Name = "Adam"
	user.Email = "adam@test.com"
	err := HawkDB().Table("users").Where("name", "Adam").First(&user)

	if err == nil {
		t.Fatal(err)
	}

	if user.Name != "Adam" {
		t.Fatalf("expected Adam, got %s", user.Name)
	}

	if user.Email != "adam@test.com" {
		t.Fatalf("expected adam@test.com, got %s", user.Email)
	}
}