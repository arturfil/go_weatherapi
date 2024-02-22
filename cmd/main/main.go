package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arturfil/go_weatherapi/router"
)

var server http.Handler

//  Right now it is very learn but we could add more fields like 
// Models, other secretkeys etc
type Application struct {
    Config Config
}

// Config - Sub struct inside Application for org. purposes
type Config struct {
    Port string
}

// Server - method that will handle the port setup
func (app *Application) Serve() error {
    server := &http.Server{
        Addr: fmt.Sprintf(":%s", app.Config.Port),
        Handler: router.Routes(), // routes added here through the http.Serer obj
    }

    return server.ListenAndServe()
}

// Create here the app object and execute it
func main() {
    var cfg Config
    port := os.Getenv("PORT")
    log.Println("API listenting on port:", port)

    cfg.Port = port

    var app = &Application{
        Config: cfg,
    }

    err := app.Serve()
    if err != nil {
        log.Fatal(err)
    }
    
}
