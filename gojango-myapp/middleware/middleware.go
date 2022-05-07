package middleware

import (
	"github.com/axsaucedo/gojango"
	"github.com/axsaucedo/gojango-myapp/data"
)

type Middleware struct {
	App    *gojango.Gojango
	Models data.Models
}
