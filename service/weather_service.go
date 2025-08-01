package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/henriquedessen/weather-app/model"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func FetchWeatherByCity(city, state string) (*model.WeatherResult, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("missing WEATHER_API_KEY")
	}

	query := url.QueryEscape(fmt.Sprintf("%s,%s", city, state))
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, query)

	fmt.Println("[WeatherAPI Request]", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weatherapi failed: %s", resp.Status)
	}

	var weather WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	c := weather.Current.TempC
	return &model.WeatherResult{
		TempC: c,
		TempF: c*1.8 + 32,
		TempK: c + 273,
	}, nil
}
