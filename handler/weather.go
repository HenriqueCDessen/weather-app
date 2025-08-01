package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/henriquedessen/weather-app/service"

	"github.com/gorilla/mux"
)

func GetWeatherByCEP(w http.ResponseWriter, r *http.Request) {
	cep := mux.Vars(r)["cep"]
	if !regexp.MustCompile(`^\d{8}$`).MatchString(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := service.FetchLocationByCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := service.FetchWeatherByCity(location.Localidade, location.UF)
	if err != nil {
		http.Error(w, "weather fetch failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
