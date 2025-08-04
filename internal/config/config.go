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
}

func NewConfig() (*Config, error) {

	addr := os.Getenv("ADDR")
	DBProtocol := os.Getenv("DB_PROTOCOL")
	DBURL := os.Getenv("DBURL")

	if addr == "" {
		return nil, errors.New("'ADDR' environment var is not set")
	}

	if DBProtocol == "" {
		return nil, errors.New("'DB_PROTOCOL' environment var is not set")
	}

	if DBURL == "" {
		return nil, errors.New("'DB_URL' environment var is not set")
	}

	return &Config{
		Addr:       addr,
		DBProtocol: DBProtocol,
		DBURL:      DBURL,
	}, nil
}
