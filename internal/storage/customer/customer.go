// Package customerstore exposes database manipulation methods for customers.
package customerstore

import (
	"database/sql"
)

type CustomerStorage struct {
	db *sql.DB
}

func NewCustomerStorage(db *sql.DB) *CustomerStorage {
	return &CustomerStorage{
		db: db,
	}
}
