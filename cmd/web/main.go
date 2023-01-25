package main

import (
	"fmt"
	"log"
	"myapp/cmd/pkg/config"
	"myapp/cmd/pkg/handlers"
	"myapp/cmd/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port:", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
