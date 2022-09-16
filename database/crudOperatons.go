package database

import (
	"database/sql"
)

type FoodItem struct {
	ID       int
	Name     string
	Category string
	Rating   int
}

//Does a Select all command to get all the info from the database
func GetFoodItem(database *sql.DB) {
	results, err := database.Query("SELECT * FROM my_industryDB.FoodItem")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var foodItem FoodItem
		err = results.Scan(&foodItem.ID, &foodItem.Name, &foodItem.Category, &foodItem.Rating)

		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(foodItem.ID, foodItem.Name, foodItem.Category, foodItem.Rating)
	}
}
