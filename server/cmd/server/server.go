package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// APIServer configuration
type APIServer struct {
	Port   int
	router *mux.Router
	server *http.Server
}

// StartServer - to start api server
func (s *APIServer) StartServer() error {

	s.server = &http.Server{
		Handler:      s.router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return s.server.ListenAndServe()
}

// StopServer will shutdown server
func (s *APIServer) StopServer() error {
	if s.server == nil {
		return fmt.Errorf("Server was not started")
	}
	return s.server.Shutdown(context.Background())
}
