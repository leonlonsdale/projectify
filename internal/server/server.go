// Package server creates and returns a http server
package server

import (
	"log"
	"net/http"
	"time"

	"github.com/leonlonsdale/projectify/internal/handlers"
)

// All the comments are belong to us

type Server struct {
	addr       string
	httpServer *http.Server
	router     *handlers.Handlers
}

func NewServer(addr string, router *handlers.Handlers) *Server {

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      router.Mux(),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	return &Server{
		addr:       addr,
		httpServer: httpServer,
		router:     router,
	}
}

func (s *Server) Serve() error {
	log.Printf("Server running on %s", s.addr)
	return s.httpServer.ListenAndServe()
}
