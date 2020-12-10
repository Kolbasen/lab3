package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type customError struct {
	Error string `json:"error"`
}

// JSON - custom json response wrapper
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR - custom error response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		jsonErr := customError{Error: err.Error()}
		JSON(w, statusCode, jsonErr)
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
