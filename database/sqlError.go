package database

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func MySqlErrorFormat(sqlErrr error) error{
	var mysqlErr *mysql.MySQLError
	var err error
	if errors.As(sqlErrr, &mysqlErr){
		switch mysqlErr.Number {
		case 1062:
			err = fmt.Errorf("Duplicate key")
			fmt.Println(err)
		case 1452:
			err = fmt.Errorf("Foreign key contraint failed")
			fmt.Println(err)
		case 1048:
			err = fmt.Errorf("Column cannot be long")
			fmt.Println(err)
		case 1406:
			err = fmt.Errorf("Data too long")
			fmt.Println(err)
		}
	}
	return err
}