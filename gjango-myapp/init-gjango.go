package main

import (
	"log"
	"os"

	"github.com/axsaucedo/gjango"
	"github.com/axsaucedo/gjango-myapp/handlers"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	g := &gjango.Gjango{}
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

	return app
}
