package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Retrieve API key from environment variables
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY not found in .env file")
		return
	}

	fmt.Println("Welcome to the weather app.....")

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com./v1/forecast.json?key=%s&q=Noida&days=1&aqi=no&alerts=no", apiKey))

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not avaliable ")
	}

	body, err := io.ReadAll((res.Body))
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}
	fmt.Println(weather)
}
