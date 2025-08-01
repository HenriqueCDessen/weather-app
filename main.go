package main

import (
	"log"
	"net/http"

	"github.com/henriquedessen/weather-app/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handler.GetWeatherByCEP).Methods("GET")

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
