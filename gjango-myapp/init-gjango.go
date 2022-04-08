package main

import (
	"log"
	"os"

	"github.com/axsaucedo/gjango"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Initialise celeritas
	g := &gjango.Gjango{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}

	g.AppName = "myapp"
	g.Debug = true

	app := &application{
		App: g,
	}

	return app
}
