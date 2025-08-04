// Package storage provides storage and database logic for crud operations, for the various domains
package storage

import (
	"database/sql"

	"github.com/leonlonsdale/projectify/internal/storage/customerstore"
)

type Storage struct {
	Customers *customerstore.CustomerStorage
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Customers: customerstore.NewCustomerStorage(db),
	}
}
