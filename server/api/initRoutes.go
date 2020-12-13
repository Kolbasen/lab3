package api

import (
	"github.com/Kolbasen/lab3/server/api/dishes"
	"github.com/Kolbasen/lab3/server/config"
	"github.com/gorilla/mux"
)

// InitRouter - init all routes in API
func InitRouter(store *dishes.Store, config *config.Config) *mux.Router {
	router := mux.NewRouter()

	dishes.AddDishesRoutes(store, config, router)

	return router
}
