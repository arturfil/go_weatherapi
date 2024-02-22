package router

import (
	"net/http"

	"github.com/arturfil/go_weatherapi/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
    router := chi.NewRouter()

    router.Use(middleware.Recoverer)
    router.Use(cors.Handler(cors.Options{ 
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
    }))

    router.Get("/api/healthcheck", handlers.HealthCheck)
    router.Get("/api/get-weather", handlers.GetWeather)

    return router
}
