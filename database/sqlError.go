package database

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func MySqlErrorFormat(sqlErr error) error {
	if sqlErr == nil {
		return nil
	}

	var mysqlErr *mysql.MySQLError

	if errors.As(sqlErr, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			return fmt.Errorf("duplicate key: %w", sqlErr)

		case 1452:
			return fmt.Errorf("foreign key constraint failed: %w", sqlErr)

		case 1048:
			return fmt.Errorf("column cannot be null: %w", sqlErr)

		case 1406:
			return fmt.Errorf("data too long: %w", sqlErr)
		}
	}

	// Important: don't swallow unknown/database errors.
	return sqlErr
}