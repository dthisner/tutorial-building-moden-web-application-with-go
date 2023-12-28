package main

import (
	"log"
	"myApp/pkg/config"
	"myApp/pkg/handlers"
	"myApp/pkg/render"
	"net/http"
)

var port = ":8081"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	log.Println("Starting server on port", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("not able to start server:", err)
	}
	// err = http.ListenAndServe(port, nil)
	// if err != nil {
	// 	log.Fatal("not able to start server:", err)
	// }
}
