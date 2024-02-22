package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/arturfil/go_weatherapi/services"
	// "github.com/go-chi/chi/v5"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	// get param queries
	lon := r.URL.Query().Get("lon") // chi.URLParam(r, "lon")
	lat := r.URL.Query().Get("lat") // chi.URLParam(r, "lat")

	apiKey := os.Getenv("SECRET_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=imperial", lat, lon, apiKey)

	// make the http request
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

    // read the response from the http get request
	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read the response body", http.StatusInternalServerError)
		return
	}

    // get the response from a json to a struct so we can 
    // manipulate the data
	var result services.WeatherResponse
	json.Unmarshal([]byte(body), &result)

    // WeatherSummary - created summary struct
	weatherSummary := &services.WeatherSummary{
		FeelsLike: temperaturefeelslike(result.Main.Temp),
		Condition: result.Weather[0].Description,
		Temp:      result.Main.Temp,
		Country:   result.Name,
		PlaceName: result.Sys.Country,
	}

    // Change the WeatherSummary from struct to json
	resJson, err := json.Marshal(&weatherSummary)
	if err != nil {
		log.Println("error:", err)
		return
	}

	// Write the response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! The app is running!")
}

// helper function to return commentary on temperature 
// (I'm from Florida 🤓 so I think it's cold after 64 F)
func temperaturefeelslike(temp float64) string {
	switch true {
	case temp >= 75:
		return "It's Hot"
	case temp >= 64 && temp < 75:
		return "It's Cool"
	case temp < 64:
		return "It's Cold"
	default:
		return "Not sure what the weather is"
	}
}
