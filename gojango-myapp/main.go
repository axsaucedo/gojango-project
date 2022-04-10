package main

import (
	"github.com/axsaucedo/gojango"
	"github.com/axsaucedo/gojango-myapp/handlers"
)

type application struct {
	App      *gojango.Gojango
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
