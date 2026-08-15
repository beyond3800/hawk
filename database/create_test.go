package database

import (
	"testing"
)


func TestCreate(t *testing.T) {

	setupTestDB(t)
	user := map[string]any{
		"name":  "Adam",
		"email": "adam@test.com",
	}
    _,err := HawkDB().Table("users").Create(user)

    if err != nil {
        t.Fatal(err)
    }

	var firstUser struct {
		Name  string
		Email string
	}
    err = HawkDB().Table("users").Where("email", user["email"]).First(&firstUser)

    if err != nil {
        t.Fatal(err)
    }
	if user["name"] != "Adam" {
		t.Fatalf("expected Adam, got %s", user["name"])
	}

	if user["email"] != "adam@test.com" {
		t.Fatalf("expected adam@test.com, got %s", user["email"])
	}


}