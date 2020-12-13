package dishes

import (
	"encoding/json"
	"net/http"

	"github.com/Kolbasen/lab3/server/tools"
	"github.com/gorilla/mux"
)

type DishesRouter = *mux.Router

type DishHandlers struct {
	store *Store
}

// AddDishesRoutes - add routes
func AddDishesRoutes(store *Store) *mux.Router {
	dishHandlers := &DishHandlers{
		store: store,
	}

	router := mux.NewRouter()

	router.HandleFunc("/dishes/list", dishHandlers.handleListDishes).Methods(http.MethodGet)
	router.HandleFunc("/dishes/order/create", dishHandlers.handleDishesOrderCreate).Methods(http.MethodPost)
	return router
}

func (d *DishHandlers) handleListDishes(w http.ResponseWriter, r *http.Request) {
	res, err := d.store.ListDishes()

	if err != nil {
		tools.ERROR(w, http.StatusInternalServerError, err)
	}

	tools.JSON(w, http.StatusOK, res)
}

type OrderCreateRequestPayload struct {
	TableID int64   `json:"table_id"`
	DishIDS []int64 `json:"dish_ids"`
}

type OrderCreateResponse struct {
	TableID         int64   `json:"table_id"`
	TotalPrice      int64   `json:"total_price"`
	TotalPriceNoTax float64 `json:"total_price_no_tax"`
	RecommendedTips float64 `json:"recommended_tips"`
}

func (d *DishHandlers) handleDishesOrderCreate(w http.ResponseWriter, r *http.Request) {
	var requesPayload OrderCreateRequestPayload
	var totalPrice int64

	if err := json.NewDecoder(r.Body).Decode(&requesPayload); err != nil {
		tools.ERROR(w, http.StatusBadRequest, err)
		return
	}
	dishes, err := d.store.GetDishesByIds(requesPayload.DishIDS)

	for _, dish := range dishes {
		totalPrice += dish.Price
	}

	if err != nil {
		tools.ERROR(w, http.StatusInternalServerError, err)
	}

	totalPriceNoTax := float64(totalPrice) * 0.8
	recommendedTips := totalPriceNoTax * 0.15

	response := &OrderCreateResponse{
		TableID:         requesPayload.TableID,
		TotalPrice:      totalPrice,
		TotalPriceNoTax: totalPriceNoTax,
		RecommendedTips: recommendedTips,
	}

	tools.JSON(w, http.StatusOK, response)
}
