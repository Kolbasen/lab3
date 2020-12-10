//+build wireinject

package main

import (
	"github.com/Kolbasen/lab3/server/api/dishes"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func ComposeApiServer() (*mux.Router, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		// NewDbConnection,
		dishes.Providers,
		// wire.Struct(new(ChatApiServer), "Port", "ChannelsHandler"),
	)
	return nil, nil
}
