package util

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func DulicatedKey(err error) bool {
	
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		fmt.Print(err)
		return mysqlErr.Number == 1062
	}
	return false
}