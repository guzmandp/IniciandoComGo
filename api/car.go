package api

import (
	"encoding/json"
	"gostart/car"
	"gostart/database"
	"net/http"
)

type CarAPI struct {
	CarDB *database.CarDB
}

func NewCarAPI(carDB *database.CarDB) *CarAPI {
	return &CarAPI{
		CarDB: carDB,
	}
}

func (api CarAPI) InsertHandler(w http.ResponseWriter, r *http.Request) {
	var car car.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = api.CarDB.Insert(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
