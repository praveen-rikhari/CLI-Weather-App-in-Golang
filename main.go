package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
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

	q := "Noida"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	fmt.Println("Welcome to the weather app.....")

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com./v1/forecast.json?key=%s&q="+q+"&days=1&aqi=no&alerts=no", apiKey))

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
	// fmt.Println(weather)

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Print("\n")
	fmt.Printf(
		"Region: %s, %s\nCountry: %s\nCurrent Temperature: %0.f°C \nCurrent Condition: %s\n",
		location.Name,
		location.Region,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
	fmt.Print("\n")
	color.Magenta("******** TODAY'S WEATHER TIMELINE ********")

	fmt.Print("\n")
	color.Green("TIME     TEMP     RAIN%    CONDITION")

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf(
			"%s -> %0.f°C     %0.f%%      %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else if hour.ChanceOfRain > 40 && hour.ChanceOfRain < 70 {
			color.HiYellow(message)
		} else {
			color.HiRed(message)
		}
	}
}
