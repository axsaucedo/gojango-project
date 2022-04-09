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

	g := &gjango.Gjango{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}

	g.AppName = "myapp"

	g.InfoLog.Println("Debug is set to", g.Debug)

	app := &application{
		App: g,
	}

	return app
}
