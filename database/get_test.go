package database

import (
	"testing"
)

func TestGet(t *testing.T) {
	setupTestDB(t)
	seedUser(t,10)
	type User struct {
		Name  string
		Email string
	}
	var user []User

	err := HawkDB().Table("users").Get(&user)
	if err != nil {
		t.Fatal(err)
	}
	if len(user) == 0{
		t.Fatal("User length should be more than zero")
	}

}