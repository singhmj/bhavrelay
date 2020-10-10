package main

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

// GetBhavData : Fetches bhav data from cache
var GetBhavData = func(w http.ResponseWriter, r *http.Request) {
	// check if symbol is valid
	symbol := mux.Vars(r)["symbol"]
	if !isSymbolValid(symbol) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid symbol in the url"))
		return
	}

	// check if series is valid
	series := mux.Vars(r)["series"]
	if !isSeriesValid(series) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid series in the url"))
		return
	}

	// fetch data from the cache and return it to the user
	data := getBhavFromCache(symbol, series)
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		log.Printf("error occurred while converting bhav data into json, err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func isSymbolValid(symbol string) bool {
	// you can add some other logic here
	return symbol != ""
}

func isSeriesValid(series string) bool {
	// you can add some other logic here
	return series != ""
}
