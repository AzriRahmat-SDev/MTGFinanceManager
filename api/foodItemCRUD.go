package api

import (
	"database/sql"
	"fmt"
	"log"
)

type FoodInfo struct {
	ID       string `json:"Id"`
	Name     string `json:"Name"`
	Category string `json:"Category"`
	Rating   string `json:"Rating"`
}

var FoodItems map[string]FoodInfo
var CardItems map[string]CardInfo

func OpenFoodItemDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_industryDB")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func InsertFoodItem(db *sql.DB, f FoodInfo) {
	query := fmt.Sprintf("INSERT INTO FoodItem (ID, Name, Category, Rating) VALUES ('%s', '%s', '%s','%s')", f.ID, f.Name, f.Category, f.Rating)
	_, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n Successful insert Food Item @ '%s'", f)
	}
}

func EditFoodItem(db *sql.DB, Name string, newFoodName string) {
	query := fmt.Sprintf("UPDATE FoodItem SET Name='%s' WHERE Name='%s'", newFoodName, Name)
	_, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n Successful update food name '%s','%s'", Name, newFoodName)
	}
}

func DeleteFoodItem(db *sql.DB, Name string) {
	query := fmt.Sprintf("DELETE FROM FoodItem WHERE Name='%s'", Name)
	_, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\nSuccessful delete Food name @ '%s'", Name)

	}
}
