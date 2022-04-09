package main

import (
	"github.com/axsaucedo/gjango"
	"github.com/axsaucedo/gjango-myapp/handlers"
)

type application struct {
	App      *gjango.Gjango
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
