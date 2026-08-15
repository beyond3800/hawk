package database


import (
	"testing"
)

func TestPaginate(t *testing.T) {
	setupTestDB(t)
	seedUser(t, 100)
	type User struct{
		Name string
		Email string
	}
	var user []User
	paginated,err := HawkDB().Table("users").Paginate(1, 10, &user)
	if err != nil {
		t.Fatal(err)
	}
	if paginated.Meta.Total != 100{
		t.Fatal("Total page is expected to be 100")
	}
	if paginated.Meta.PerPage != 10{
		t.Fatal("Total page is expected to be 10")
	}

}