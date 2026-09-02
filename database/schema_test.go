package database

import (
	"testing"
)

func TestTableHasTwoPrimaryKey(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.BigInt("user").Primary().AutoIncrement()
	table.String("name",100)

	
	if err := table.Create(); err == nil{
		t.Fatal("An error is expected because two primary key is not allowed",err)
	}
}
func TestTableHasPrimaryKey(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.String("name",100)

	
	if err := table.Create(); err != nil{
		t.Fatal(err)
	}
}
func TestStringHasPrimaryKey(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.String("user", 255).Primary().AutoIncrement()
	table.String("name",100)

	
	if err := table.Create(); err == nil{
		t.Fatal(err)
	}
}
func TestComposedPrimary(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("user_id")
	table.BigInt("post_id")
	table.ComposedPrimary("user_id","post_id")

	
	if err := table.Create(); err == nil{
		t.Fatal(err)
	}
}
func TestTableHasComposedAndPrimaryKey(t *testing.T){
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.BigInt("user_id")
	table.BigInt("post_id")
	table.ComposedPrimary("user_id","post_id")

	
	if err := table.Create(); err == nil{
		t.Fatal(err)
	}
}
func TestHasSameIndexName(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.String("content",255).FullText("content_fullText")
	table.String("comment",255)
	table.BigInt("user_id")
	table.BigInt("post_id")
	table.FullText("content_fullText", "comment")
	
	if err := table.Create(); err == nil{
		t.Fatal(err)
	}
}
func TestColumnDoesNotExist(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.String("content",255).Index("content_full_Text")
	table.String("poster",255).Index("content_full_Texts")
	table.String("comment",255)
	table.BigInt("user_id")
	table.BigInt("post_id")
	table.Index("comment_fullText", "comment")
	// table.FullText("content_full_Text", "comment")
	// table.FullText("comment_fullText", "comment")
	
	if err := table.Create(); err == nil{
		t.Fatal(err)
	}
}

func TestForeignIDQuery(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.String("content",255)
	table.String("poster",255)
	table.String("comment",255)
	table.ForeignID("user_id").Constrained("users").OnDelete().SetNull()
	// table.ForeignID("user_id").Constrained("users")
	table.FullText("user_id", "comment")
	// table.ForeignID("user_id").Constrained("admins")
	table.BigInt("post_id")
	table.Index("comment_fullText", "comment")
	
	if err := table.Create(); err != nil{
		t.Fatal(err)
	}
}
func TestEnumEmptyValue(t *testing.T){
	setupTestDB(t)
	table := HawkDB().Schema().Table("test1")
	table.BigInt("id").Primary().AutoIncrement()
	table.Enum("state","rejected","accepted")
	if err := table.Create(); err == nil{
		t.Fatal(err)
	}
}
func TestAlterColumn(t *testing.T) {
	setupTestDB(t)
	table := HawkDB().Schema().Table("users")
	if err := table.AlterColumn("name").Rename("username").Execute(); err != nil{
		t.Fatal(err)
	}

	
}