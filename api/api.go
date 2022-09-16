package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const baseURL = "http://localhost:2020/api/v1"
const scryFallURL = "http://api.scryfall.com/cards/named?"

func StartServer() {
	//initiate the making of the map
	FoodItems = make(map[string]FoodInfo)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/food", ListOfFoodItems)

	//this is engaged by a form method in the in the client to do REST operations
	// router.HandleFunc("/api/v1/food/", GetAllFoodItems).Methods("GET")
	// router.HandleFunc("/api/v1/food/{foodName}", FoodItemHandler).Methods("GET", "POST", "PUT", "DELETE")

	router.HandleFunc("/api/v1/cardName/", GetAllCardItemIntoMap).Methods("GET")
	router.HandleFunc("/api/v1/cardName/{cardName}", CardHandler).Methods("GET", "POST", "PUT", "DELETE")

	log.Fatal(http.ListenAndServe(":2020", router))
}
