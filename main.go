package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/texto/diario", GetMonth).Methods("GET")
	router.HandleFunc("/texto/diario/{id}", GetDay).Methods("GET")
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
