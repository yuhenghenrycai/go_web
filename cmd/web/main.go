package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yuhenghenrycai/go_web/pkg/config"
	"github.com/yuhenghenrycai/go_web/pkg/handlers"
	"github.com/yuhenghenrycai/go_web/pkg/render"
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

	render.NewRender(&app)

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	fmt.Printf("starting web app on portnumber %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	srv.ListenAndServe()
}
