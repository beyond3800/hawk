package util

// import (
// 	"fmt"

// 	"github.com/beyond3800/hawk/db"
// 	"github.com/google/uuid"
// )

// func UniqueId(table string) string {
// 	id:= uuid.NewString()
	
// 	query := fmt.Sprintf("SELECT id FROM %s Where id= ?", table)
// 	err:= db.DB.QueryRow(query,id).Scan()
	
// 	if err != nil {
// 		return id
// 	}else{
// 		UniqueId(table)
// 		return id
// 	}
// }

