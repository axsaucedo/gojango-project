package main

import (
	"log"
	"os"

	"github.com/axsaucedo/gojango"
	"github.com/axsaucedo/gojango-myapp/data"
	"github.com/axsaucedo/gojango-myapp/handlers"
	"github.com/axsaucedo/gojango-myapp/middleware"
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

	myMiddleware := &middleware.Middleware{
		App: g,
	}

	myHandlers := &handlers.Handlers{
		App: g,
	}

	app := &application{
		App:        g,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
