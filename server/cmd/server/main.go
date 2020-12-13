package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Kolbasen/lab3/server/config"
	"github.com/Kolbasen/lab3/server/db"
)

// NewDbConnection - func to establish DB connection
func NewDbConnection(config *config.Config) (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     config.Db.DbName,
		User:       config.Db.User,
		Host:       config.Db.Host,
		DisableSSL: config.Db.DisableSSL,
	}
	return conn.Open()
}

func main() {
	flag.Parse()

	if server, err := ComposeApiServer(8080, "./config/dev-config.json"); err == nil {

		go func() {
			log.Println("Starting chat server...")

			err := server.StartServer()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.StopServer(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize chat server: %s", err)
	}
}
