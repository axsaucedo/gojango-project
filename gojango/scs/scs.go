package scs

import (
    "net/http"
    "time"
)

type Cookie struct {
	Persist  bool
	Name     string
	Secure   bool
	Domain   string
	SameSite http.SameSite
}

type SessionManager struct {
	Lifetime time.Duration
	Cookie   Cookie
}

func New() *SessionManager {
	return &SessionManager{}
}

func (s *SessionManager) LoadAndSave(h http.Handler) http.Handler {
    return h
}
