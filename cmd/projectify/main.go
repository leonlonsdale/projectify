// Package main is the entry point for the application
package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/leonlonsdale/projectify/internal/auth"
	"github.com/leonlonsdale/projectify/internal/config"
	"github.com/leonlonsdale/projectify/internal/database"
	"github.com/leonlonsdale/projectify/internal/router"
	"github.com/leonlonsdale/projectify/internal/server"
	"github.com/leonlonsdale/projectify/internal/storage"
)

type Application struct {
	cfg    *config.Config
	auth   *auth.Auth
	store  *storage.Storage
	routes *router.Router
	db     *sql.DB
	server *server.Server
}

func (a *Application) setup() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	a.cfg = cfg

	a.auth = auth.NewAuth(a.cfg)
	db, err := database.NewDB(a.cfg)
	if err != nil {
		return fmt.Errorf("init db: %w", err)
	}

	a.db = db
	a.store = storage.NewStorage(a.db)
	a.routes = router.NewRouter(a.store, a.cfg, a.auth)
	a.server = server.NewServer(a.cfg.Addr, a.routes)

	return nil
}

func main() {
	app := &Application{}

	if err := app.setup(); err != nil {
		slog.Error("could not initialise the application", "error", err)
		os.Exit(1)
	}

	if err := app.server.Serve(); err != nil {
		slog.Error("could not start the server", "error", err)
		os.Exit(1)
	}
}
