//+build wireinject

package main

import (
	"github.com/Kolbasen/lab3/server/api/dishes"
	"github.com/google/wire"
)

func ComposeApiServer(port int) (*APIServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		// NewDbConnection,
		dishes.Providers,
		wire.Struct(new(APIServer), "Port", "router"),
	)
	return nil, nil
}
