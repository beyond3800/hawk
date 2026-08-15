package database

import (
	"errors"
	"testing"
)


func TestTransaction(t *testing.T){
	setupTestDB(t)
	createWalletTable(t)
	err := HawkDB().Transaction(func(tx *Builder) error {

		result, err := tx.Table("users").Create(map[string]any{
			"name": "Adams",
			"email": "adams@gmail.com",
		})
		if err != nil {
			t.Fatal("Unable to create user", err)
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			t.Fatal("Unable to get last insert ID", err)
			return err
		}
		_, err = tx.Table("wallets").Create(map[string]any{
			"user_id": id,
			"balance": 0,
		})
		if err != nil {
			t.Fatal("Unable to create wallet", err)
			return err
		}
		return nil
	})
	assertCount(t, "users", 1)
	assertCount(t, "wallets", 1)
	if err != nil {
		t.Fatal("Transaction failed", err)
	}
}
func TestTransactionRollback(t *testing.T) {

	setupTestDB(t)
	createWalletTable(t)

	err := HawkDB().Transaction(func(tx *Builder) error {

		_, err := tx.Table("users").Create(map[string]any{
			"name": "Adam",
			"email": "adam@test.com",
		})

		if err != nil {
			return err
		}

		return errors.New("force rollback")
	})

	if err == nil {
		t.Fatal("expected rollback error")
	}

	assertCount(t, "users", 0)
	assertCount(t, "wallets", 0)
}
func TestTransactionRollbackOnInsertError(t *testing.T) {
	setupTestDB(t)
	createWalletTable(t)

	err := HawkDB().Transaction(func(tx *Builder) error {

		_, err := tx.Table("users").Create(map[string]any{
			"name":  "Adam",
			"email": "adam@test.com",
		})
		if err != nil {
			return err
		}

		// Intentionally use a column that does not exist.
		_, err = tx.Table("wallets").Create(map[string]any{
			"user_id":         1,
			"balance":         0,
			"invalid_column":  "this should fail",
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		t.Fatal("expected transaction to fail")
	}

	assertCount(t, "users", 0)
	assertCount(t, "wallets", 0)
}
