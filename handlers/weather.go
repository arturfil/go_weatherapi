package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "github.com/go-chi/chi/v5"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
    lon := r.URL.Query().Get("lon") // chi.URLParam(r, "lon")
    lat := r.URL.Query().Get("lat") // chi.URLParam(r, "lat")

    apiKey := os.Getenv("SECRET_KEY")
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, apiKey) 

    response, err := http.Get(url)
    if err != nil {
        http.Error(w, "Error", http.StatusInternalServerError)
        return 
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
		http.Error(w, "Failed to read the response body", http.StatusInternalServerError)
		return
	}

	// Write the response to the client
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Response from %s:\n%s", url, body)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello! The app is running!")
}

