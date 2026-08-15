package database


import (
	"testing"
)

func TestUpdate(t *testing.T) {
	setupTestDB(t)

	seedUser(t, 10)
	result, err := HawkDB().Table("users").Where("id", 1).Update(map[string]any{
		"email": "",
	})
	if err != nil {
		t.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if rowsAffected != 1 {
		t.Fatalf("expected 1 row affected, got %d", rowsAffected)
	}
}