package handlers

import (
	"fmt"
	"net/http"

	"github.com/axsaucedo/gjango"
)

type Handlers struct {
	App *gjango.Gjango
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	fmt.Errorf("error", err)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
