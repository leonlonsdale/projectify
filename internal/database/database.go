// Package database exposes a function NewDB that returns a new pg database connection
package database

import (
	"database/sql"

	"github.com/leonlonsdale/projectify/internal/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBProtocol, cfg.DBURL)
	if err != nil {
		// TODO: Handle error
		return nil, nil
	}

	if err := db.Ping(); err != nil {
		// TODO: Handle error
		return nil, nil
	}

	return db, nil
}
