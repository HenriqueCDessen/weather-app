package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henriquedessen/weather-app/handler"

	"github.com/gorilla/mux"
)

func TestGetWeatherByCEP_InvalidCEP(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather/123", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handler.GetWeatherByCEP)
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, got %d", rr.Code)
	}
}

func TestGetWeatherByCEP_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather/00000000", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handler.GetWeatherByCEP)
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", rr.Code)
	}
}
