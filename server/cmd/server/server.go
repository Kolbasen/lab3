package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hell")
	// router := mux.NewRouter()
	// dishes.AddDishesRoutes(router)
	router, err := ComposeApiServer()

	if err != nil {
		return
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}
