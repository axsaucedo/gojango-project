package main

import (
	"log"
	"os"

	"github.com/axsaucedo/gojango"
	"github.com/axsaucedo/gojango-myapp/data"
	"github.com/axsaucedo/gojango-myapp/handlers"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	g := &gojango.Gojango{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}

	g.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: g,
	}

	app := &application{
		App:      g,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	return app
}
