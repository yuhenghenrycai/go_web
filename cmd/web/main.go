package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/yuhenghenrycai/go_web/pkg/config"
	"github.com/yuhenghenrycai/go_web/pkg/handlers"
	"github.com/yuhenghenrycai/go_web/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 + time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
