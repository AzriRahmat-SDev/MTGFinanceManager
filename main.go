package main

import (
	"GoIndustry/api"
	"GoIndustry/database"
	"GoIndustry/functions"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_industryDB")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Database successfully opened")

	database.GetFoodItem(db)

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
