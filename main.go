package main

import (
	"database/sql"
	"gostart/car"
	"gostart/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./cars.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ferrariDB := database.NewCarDB(db)

	ferrari := car.Car{
		Brand: "Ferrari",
		Plate: "488 GTB",
		Color: "red",
		Year:  2021,
	}

	err = ferrariDB.Insert(ferrari)
	if err != nil {
		panic(err)
	}
}
