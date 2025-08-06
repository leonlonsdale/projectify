// Package database exposes a function NewDB that returns a new pg database connection
package database

import (
	"database/sql"

	"github.com/leonlonsdale/projectify/internal/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
