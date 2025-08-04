// Package server creates and returns a http server
package server

import (
	"log"
	"net/http"
	"time"

	"github.com/leonlonsdale/projectify/internal/router"
)

// All the comments are belong to us

type Server struct {
	addr       string
	httpServer *http.Server
	router     *router.Router
}

func NewServer(addr string, router *router.Router) *Server {

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
