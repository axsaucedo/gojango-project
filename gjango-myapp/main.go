package main

import "github.com/axsaucedo/gjango"

type application struct {
	App *gjango.Gjango
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
