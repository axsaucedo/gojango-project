package handlers

import (
	"net/http"

	"github.com/axsaucedo/gojango"
)

type Handlers struct {
	App *gojango.Gojango
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
