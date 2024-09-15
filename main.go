package main

import (
	"database/sql"
	"gostart/api"
	"gostart/database"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./cars.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	carDB := database.NewCarDB(db)
	carAPI := api.NewCarAPI(carDB)
	http.HandleFunc("POST /cars", carAPI.InsertHandler)
	http.ListenAndServe(":7171", nil)
}
