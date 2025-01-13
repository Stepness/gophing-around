package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	Hourly struct {
		Time        []string  `json:"time"`
		Temperature []float32 `json:"temperature_2m"`
	} `json:"hourly"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func main() {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=45.4643&longitude=9.1895&hourly=temperature_2m&timezone=Europe%2FBerlin&forecast_days=1")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	hours, temperatures := weather.Hourly.Time, weather.Hourly.Temperature
	for i, h := range hours {
		date, _ := time.Parse("2006-01-02T15:04", h)
		if date.Before(time.Now()) {
			continue
		}
		fmt.Printf("At %d, expect %.0f\n", date.Hour(), temperatures[i])
	}
}
