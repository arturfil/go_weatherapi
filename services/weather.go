package services

// In this package we would see the services interacting with a database and the models related
// Since we don't need a database connection, we only have the models written here.
// Adversely we could place these models in the handlers for this project,
// But in a larger project these would go here with the additional DB queries

// WeatherResponse - will have all the information from the api
// I decided to separate the structures because if we want to change this
// sturct later, it will be very flexible
type WeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Sys     Sys       `json:"sys"`
	Name    string    `json:"name"`
}

// weather sturct with necessary attributes
// We only really need - "Description"
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main weather struct from the response
// we could use all the fields but we only need - "Temp"
type Main struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Temp_min   float64 `json:"temp_min"`
	Temp_max   float64 `json:"temp_max"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
	SeaLevel   int     `json:"sea_level"`
	GrndLevel  int     `json:"grnd_level"`
}

// Sys - has more fields but we only want the country field from "Sys"
type Sys struct {
    Country string `json:"country"`
}

// WeatherSummary - will be the summarized weather reponse
type WeatherSummary struct {
	FeelsLike string  `json:"feels_like"`
	Condition string  `json:"conditon"`
	Temp      float64 `json:"temp"`
	PlaceName string  `json:"place_name"`
	Country   string  `json:"country"`
}
