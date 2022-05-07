package main

import (
	"github.com/axsaucedo/gojango"
	"github.com/axsaucedo/gojango-myapp/data"
	"github.com/axsaucedo/gojango-myapp/handlers"
	"github.com/axsaucedo/gojango-myapp/middleware"
)

type application struct {
	App        *gojango.Gojango
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
