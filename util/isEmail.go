package util

// import (
// 	"database/sql"
// 	"fmt"
// 	"strings"

// 	"github.com/beyond3800/hawk/db"
// )

// func IsEmail(email string) bool{
// 	com := strings.HasSuffix(email, ".com")
// 	at := strings.Contains(email,"@")
	
// 	if com && at {
// 		return true
// 	}
// 	return false
// }

// func IsEmailAvailable(table string, email string) bool{
// 	query := fmt.Sprintf("SELECT id FROM %s WHERE email=?",table)

// 	var id string
// 	err := db.DB.QueryRow(query,email).Scan(&id)
// 	if err == sql.ErrNoRows{
// 		return false
// 	}
// 	return true
// }