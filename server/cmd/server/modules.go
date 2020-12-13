//+build wireinject

package main

import (
	"github.com/Kolbasen/lab3/server/api"
	"github.com/Kolbasen/lab3/server/api/dishes"
	"github.com/Kolbasen/lab3/server/config"
	"github.com/google/wire"
)

func ComposeApiServer(port int, configFile string) (*APIServer, error) {
	wire.Build(
		config.LoadConfiguration,
		NewDbConnection,
		dishes.Providers,
		api.InitRouter,
		wire.Struct(new(APIServer), "Port", "router"),
	)
	return nil, nil
}
