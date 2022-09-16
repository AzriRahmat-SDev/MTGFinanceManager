package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type FoodMap map[string]FoodInfo

func ListOfFoodItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of all food items \n")

	kv := r.URL.Query()
	for k, v := range kv {
		fmt.Println(k, v)
	}

	json.NewEncoder(w).Encode(FoodItems)
	fmt.Println(FoodItems)
}

func GetAllFoodItems(w http.ResponseWriter, r *http.Request) {
	fMap := makeFoodItemMap("")
	json.NewEncoder(w).Encode(fMap)
}

//this concept will allow me to insert a var that uses a map to search a KeyValue pair
func FoodItemHandler(w http.ResponseWriter, r *http.Request) {
	db := OpenFoodItemDB()
	params := mux.Vars(r)
	fmt.Fprint(w, "Details of food item: "+params["foodName"])
	fmt.Fprint(w, "\n"+r.Method)

	if r.Method == "GET" {
		if _, ok := FoodItems[params["foodName"]]; ok {
			json.NewEncoder(w).Encode(FoodItems[params["foodName"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("\n404 - Not Found"))
		}
	}

	if r.Method == "POST" {
		var newFoodItem FoodInfo
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newFoodItem)

			//.Name checks for the json if \"Name\": entity is in the json body
			if newFoodItem.Name == "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("\n422 - Please supply the name of the food item"))
				return
			}

			//scanning through the map and see if the name exist
			//is !ok then add into the map, !ok means name given is not in the map
			//if there is a match then proceed to throw a duplicate msg
			if _, ok := FoodItems[params["foodName"]]; !ok {
				// FoodItems[params["foodName"]] = newFoodItem
				InsertFoodItem(db, newFoodItem)
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("\n201 - Food Item was added into the DB: " + params["foodName"]))
			} else {
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte("\n409 - Duplicate Food name"))
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("\n422 - Please supply the name of the food item"))
		}
	}

	if r.Method == "DELETE" {
		if foodDBRowExists(params["foodName"], "Name") {
			DeleteFoodItem(db, params["foodName"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("\n201 - Food item was deleted"))
		} else {
			fmt.Println(params["foodName"])
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("\n404 - Food name given was not Found"))
		}
	}

	if r.Method == "PUT" {
		var newFoodItem FoodInfo

		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			json.Unmarshal(reqBody, &newFoodItem)

			if newFoodItem.Name == "" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("\n422 - Please supply the name of the food item"))
				return
			}
			//checking if the name has a duplicate
			if _, ok := FoodItems[params["foodName"]]; ok {
				InsertFoodItem(db, newFoodItem)
				fmt.Println("Food item has been updated")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("201 - Food item added: " + params["foodName"]))
			} else {
				if newFoodItem.Name != "" {
					EditFoodItem(db, params["foodName"], newFoodItem.Name)
				}
				w.Write([]byte("201-" + params["foodName"] + " has been updated." + newFoodItem.Name))
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("\n422 - Please supply the name of the food item in JSON format"))
		}
	}
}

func makeFoodItemMap(val string) FoodMap {

	foodMap := make(map[string]FoodInfo)
	Database := OpenFoodItemDB()
	defer Database.Close()

	if val == "" {
		query := fmt.Sprintf("SELECT * FROM my_industryDB.FoodItem")
		res, err := Database.Query(query)
		if err != nil {
			fmt.Printf("")
		}
		for res.Next() {
			var f FoodInfo
			res.Scan(&f.ID, &f.Name, &f.Category, &f.Rating)
			foodMap[f.Name] = f
		}
	} else {
		result, err := Database.Query("SELECT * FROM my_industryDB.FoodItem WHERE Name = '" + val + "'")
		if err != nil {
			fmt.Println(err.Error())
		}
		for result.Next() {
			var f FoodInfo
			err := result.Scan(&f.ID, &f.Name, &f.Category)
			if err != nil {
				fmt.Println(err.Error())
			}
			foodMap["FoodInfo"] = f
		}
	}

	return foodMap
}

func foodDBRowExists(val string, column string) bool {
	Database := OpenFoodItemDB()
	defer Database.Close()
	r := false
	result, err := Database.Query("SELECT EXISTS(SELECT * FROM my_industryDB.FoodItem WHERE " + column + "='" + val + "')")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err = result.Scan(&r)
		if err != nil {
			panic(err.Error())
		}
	}
	return r
}
