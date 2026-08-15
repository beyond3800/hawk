package database


import (
	"testing"
)

func TestDelete(t *testing.T) {
	setupTestDB(t)
	seedUser(t, 10)
	_, err := HawkDB().Table("users").Where("id", 1).Delete()
	if err != nil {
		t.Fatal(err)
	}
}