package db

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

// Connection - Database connectoion
type Connection struct {
	DbName         string
	User, Password string
	Host           string
	DisableSSL     bool
}

// ConnectionURL - create connection URL
func (c *Connection) ConnectionURL() string {
	dbURL := &url.URL{
		Scheme: "postgres",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbURL.RawQuery = url.Values{
			"sslmode": []string{"disable"},
		}.Encode()
	}
	return dbURL.String()
}

// Open connection for db
func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.ConnectionURL())
}
