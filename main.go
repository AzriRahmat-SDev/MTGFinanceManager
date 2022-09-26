package main

import (
	"GoIndustry/api"
	"GoIndustry/functions"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1", functions.Home)

	go func() {
		fmt.Println("Starting API")
		api.StartServer()
		wg.Done()
	}()

	go func() {
		fmt.Println("Starting server of client")
		log.Fatal(http.ListenAndServe(":8080", router))
		wg.Done()
	}()

	wg.Wait()
}
