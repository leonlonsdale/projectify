// Package config returns a configuration struct
package config

import (
	"errors"
	"os"
)

type Config struct {
	Addr      string `env:"ADDR"`
	DBURL     string `env:"DB_URL"`
	JWTSecret string `env:"JWT_SECRET"`
}

func NewConfig() (*Config, error) {

	addr := os.Getenv("ADDR")
	dbURL := os.Getenv("DB_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return nil, errors.New("'JWT_SECRET' environment var is not set")
	}

	if addr == "" {
		return nil, errors.New("'ADDR' environment var is not set")
	}

	if dbURL == "" {
		return nil, errors.New("'DB_URL' environment var is not set")
	}

	return &Config{
		Addr:      addr,
		DBURL:     dbURL,
		JWTSecret: jwtSecret,
	}, nil
}
