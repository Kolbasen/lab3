package dishes

import (
	"net/http"

	"github.com/Kolbasen/lab3/server/tools"
	"github.com/gorilla/mux"
)

// AddDishesRoutes - add routes
func AddDishesRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/dishes/list", handleListDishes).Methods(http.MethodGet)
	return router
}

func handleListDishes(w http.ResponseWriter, r *http.Request) {
	tools.JSON(w, http.StatusOK, "Dishes list")
}
