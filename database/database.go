package database

import (
	"database/sql"
	"gostart/car"
)

type CarDB struct {
	Db *sql.DB
}

func NewCarDB(db *sql.DB) *CarDB {
	return &CarDB{Db: db}
}

func (d CarDB) Insert(car car.Car) error {
	insert := "INSERT INTO cars (brand, plate, color, year) VALUES (?, ?, ?, ?)"
	_, err := d.Db.Exec(insert, car.Brand, car.Plate, car.Color, car.Year)

	if err != nil {
		return err
	}
	return nil
}
