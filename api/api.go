package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const baseURL = "http://localhost:2020/api/v1"
const scryFallURL = "http://api.scryfall.com/cards/named?"

func StartServer() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/cardName/", GetAllCardItemIntoMap).Methods("GET")
	router.HandleFunc("/api/v1/cardName/{cardName}", CardHandler).Methods("GET", "POST", "PUT", "DELETE")

	log.Fatal(http.ListenAndServe(":2020", router))
}
