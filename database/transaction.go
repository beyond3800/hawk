package database

import "fmt"

func (db *DB) Transaction(fn func(bd *Builder) error) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	builder := &Builder{
		db: db,
		tx: tx,
	}

	if err := fn(builder); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}