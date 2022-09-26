package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CardMap map[string]CardInfo

var CardItem map[string]CardInfo

var multiverse_id string

func GetNamedCardFromAPI(w http.ResponseWriter, r *http.Request) (CardInfo, error) {
	var cardInfo CardInfo

	cardMap := make(map[string]CardInfo)
	params := mux.Vars(r)

	url := "https://api.scryfall.com/cards/named?fuzzy=" + params["cardName"]
	err := GetCardJson(url, &cardInfo)

	if err != nil {
		fmt.Printf("Error getting from api: %s\n", err.Error())
	} else {
		multiverse_id = ReadSliceValue(cardInfo.Multiverse_id)
		tcgIdString := strconv.FormatInt(int64(cardInfo.TcgPlayer_id), 10)
		Multiverse_URL := "https://api.scryfall.com/cards/multiverse/" + multiverse_id
		TcgPlayer_URL := "https://api.scryfall.com/cards/tcgplayer/" + tcgIdString
		if multiverse_id != "0" {
			response := MultiverseAPICall(Multiverse_URL)
			fmt.Printf("\nNormal Non-Foil Market Price: " + response.CardPrice.PricesNormal)
			cardMap[response.Name] = cardInfo
		} else if tcgIdString != "0" {
			response := TcgPlayerAPICall(TcgPlayer_URL)
			fmt.Printf("\nNormal Non-Foil TCG Player Price: " + response.CardPrice.PricesNormal)
			cardMap[response.Name] = cardInfo
		} else {
			return cardInfo, errors.New("Please Provide a Valid card name")
		}

	}
	return cardInfo, nil
}

func GetCardJson(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

//Gets all the available data from DB into a map
func GetAllCardItemIntoMap(w http.ResponseWriter, r *http.Request) {
	fMap := makeCardMap("")
	json.NewEncoder(w).Encode(fMap)
}

func CardHandler(w http.ResponseWriter, r *http.Request) {

	db := OpenCardDB()
	defer db.Close()
	params := mux.Vars(r)
	if r.Method == "GET" {
		GetNamedCardFromAPI(w, r)
	}
	if r.Method == "DELETE" {
		if cardDBRowExists(params["cardName"], "Name") {
			DeleteCardItem(db, params["cardName"])
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("\n201 - Card item was deleted"))
		} else {
			fmt.Println(params["cardName"])
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("\n404 - Card name given was not Found"))
		}
	}
	if r.Method == "POST" {
		var newCardName CardInfo
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newCardName)
			//check with DB if row with name already exist
			if !cardDBRowExists(params["cardName"], "Name") {
				res, err := GetNamedCardFromAPI(w, r)
				if err == nil && res.Name != "" {
					InsertCardName(db, res, multiverse_id)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("\n201 - Card name was added into the DB: " + params["cardName"]))
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("\n422 - Too many cards match ambiguous name " + params["cardName"] + ". Add more words to refine your search."))
				}
			} else {
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte("\n409 - Duplicate Card name"))
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("\n422 - Please supply the name of Card"))
		}
	}

}

func makeCardMap(val string) CardMap {
	cardMap := make(map[string]CardInfo)
	Database := OpenCardDB()
	defer Database.Close()

	if val == "" {
		query := fmt.Sprintf("SELECT * FROM my_db.CardItem")

		res, err := Database.Query(query)
		if err != nil {
			fmt.Printf("")
		}

		for res.Next() {
			var f CardInfo
			res.Scan(&f.ID, &f.Name, &f.CardMarket_id, &f.Multiverse_id, &f.TcgPlayer_id, &f.MTGO_id, &f.Reserved)
			cardMap[f.Name] = f
		}

	} else {
		query := fmt.Sprintf("SELECT * FROM my_db.CardItem WHERE Name = '" + val + "'")
		res, err := Database.Query(query)
		if err != nil {
			fmt.Println(err.Error())
		}
		for res.Next() {
			var f CardInfo
			err := res.Scan(&f.ID, &f.Name, &f.CardMarket_id, &f.Multiverse_id, &f.TcgPlayer_id, &f.MTGO_id, &f.Reserved)
			if err != nil {
				fmt.Println(err.Error())
			}
			cardMap["CardInfo"] = f
		}
	}
	return cardMap
}

func cardDBRowExists(val string, column string) bool {
	Database := OpenCardDB()
	defer Database.Close()

	ifExists := false

	result, err := Database.Query("SELECT EXISTS(SELECT * FROM my_db.CardItem WHERE " + column + "='" + val + "')")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err = result.Scan(&ifExists)
		if err != nil {
			panic(err.Error())
		}
	}
	return ifExists
}

func ReadSliceValue(val []int) string {
	if val != nil {
		for _, element := range val {
			var newId string
			newId = strconv.FormatInt(int64(element), 10)
			return newId
		}
	}
	return "0"
}

func MultiverseAPICall(val string) CardInfo {
	var cardInfo CardInfo
	cardMap := make(map[string]CardInfo)
	err := GetCardJson(val, &cardInfo)
	if err != nil {
		fmt.Printf("Error getting from api: %s\n", err.Error())
	} else {
		cardMap[cardInfo.Name] = cardInfo
	}
	return cardInfo
}

func TcgPlayerAPICall(val string) CardInfo {
	var cardInfo CardInfo
	cardMap := make(map[string]CardInfo)
	err := GetCardJson(val, &cardInfo)
	if err != nil {
		fmt.Printf("Error getting from api: %s\n", err.Error())
	} else {
		cardMap[cardInfo.Name] = cardInfo
	}
	return cardInfo
}
