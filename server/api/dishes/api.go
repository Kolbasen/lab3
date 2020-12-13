package dishes

import (
	"encoding/json"
	"net/http"

	"github.com/Kolbasen/lab3/server/config"
	"github.com/Kolbasen/lab3/server/tools"
	"github.com/gorilla/mux"
)

// DishHandlers - strutct for dish API
type DishHandlers struct {
	store  *Store
	config *config.Config
}

// AddDishesRoutes - add routes for dishes endpoint
func AddDishesRoutes(store *Store, config *config.Config, router *mux.Router) {
	dishHandlers := &DishHandlers{
		store:  store,
		config: config,
	}

	router.HandleFunc("/dishes/list", dishHandlers.handleListDishes).Methods(http.MethodGet)
	router.HandleFunc("/dishes/order/create", dishHandlers.handleDishesOrderCreate).Methods(http.MethodPost)
}

func (d *DishHandlers) handleListDishes(w http.ResponseWriter, r *http.Request) {
	res, err := d.store.ListDishes()

	if err != nil {
		tools.ERROR(w, http.StatusInternalServerError, err)
	}

	tools.JSON(w, http.StatusOK, res)
}

// OrderCreateRequestPayload - body in request
type OrderCreateRequestPayload struct {
	TableID int64   `json:"table_id"`
	DishIDS []int64 `json:"dish_ids"`
}

// OrderCreateResponse - response type
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

	totalPriceNoTax := float64(totalPrice) * d.config.App.TaxPercent
	recommendedTips := totalPriceNoTax * d.config.App.RecommendedTips

	response := &OrderCreateResponse{
		TableID:         requesPayload.TableID,
		TotalPrice:      totalPrice,
		TotalPriceNoTax: totalPriceNoTax,
		RecommendedTips: recommendedTips,
	}

	tools.JSON(w, http.StatusOK, response)
}
