package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMonth(w http.ResponseWriter, r *http.Request) {
	data := getFile()
	json.NewEncoder(w).Encode(data)
}

func GetDay(w http.ResponseWriter, r *http.Request) {
	data := getFile()
	params := mux.Vars(r)
	for _, item := range data.Agosto {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Month{})
}
