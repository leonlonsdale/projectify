// Package config returns a configuration struct
package config

import (
	"errors"
	"os"
)

type Config struct {
	Addr       string
	DBProtocol string
	DBURL      string
	JWTSecret  string
}

func NewConfig() (*Config, error) {

	addr := os.Getenv("ADDR")
	dbProtocol := os.Getenv("DB_PROTOCOL")
	dbURL := os.Getenv("DBURL")
	jwtSecret := os.Getenv("JWT_SECRET")

	if addr == "" {
		return nil, errors.New("'ADDR' environment var is not set")
	}

	if dbProtocol == "" {
		return nil, errors.New("'DB_PROTOCOL' environment var is not set")
	}

	if dbURL == "" {
		return nil, errors.New("'DB_URL' environment var is not set")
	}

	return &Config{
		Addr:       addr,
		DBProtocol: dbProtocol,
		DBURL:      dbURL,
		JWTSecret:  jwtSecret,
	}, nil
}
