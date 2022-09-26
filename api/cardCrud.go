package api

import (
	"database/sql"
	"fmt"
	"log"
)

type CardInfo struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	CardMarket_id int         `json:"cardmarket_id"`
	TcgPlayer_id  int         `json:"tcgplayer_id"`
	MTGO_id       int         `json:"mtgo_id"`
	Reserved      bool        `json:"reserved"`
	Multiverse_id []int       `json:"multiverse_ids"`
	CardPrice     Card_Prices `json:"prices"`
	Object        string
}

type Card_Prices struct {
	PricesNormal string `json:"usd"`
	PricesFoil   string `json:"usd_foil"`
	MTGO_Tix     string `json:"tix"`
}

const connection string = "root:password@tcp(localhost:56297)/my_db"

func OpenCardDB() *sql.DB {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

//Create and insert into DB, links with post request
func InsertCardName(db *sql.DB, f CardInfo, multiverse_ids string) {
	query := fmt.Sprintf("INSERT INTO CardItem (ID,Name,CardMarket_id, TcgPlayer_id,Multiverse_id,MTGO_id,Reserved,PriceNormal,PriceFoil,MTGO_Tix) VALUES ('%s', '%s', '%v','%v','%v','%v','%v','%s','%s','%s')", f.ID, f.Name, f.CardMarket_id, f.TcgPlayer_id, multiverse_ids, f.MTGO_id, f.Reserved, f.CardPrice.PricesNormal, f.CardPrice.PricesFoil, f.CardPrice.MTGO_Tix)
	_, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\nSuccessful insert @ '%s' into CardItem database\n", f.Name)
	}
}

func DeleteCardItem(db *sql.DB, Name string) {
	query := fmt.Sprintf("DELETE FROM CardItem WHERE Name='%s'", Name)
	_, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Successful delete @ '%s' from CardItem database", Name)
	}
}

//Read item from the DB

//Need to pull info from DB 1st put it into a Map the display hence no manipulation of the DB
