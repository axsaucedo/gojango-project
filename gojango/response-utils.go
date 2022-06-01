package gojango

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
)

func (g *Gojango) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gojango) WriteXML(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gojango) DownloadFile(w http.ResponseWriter, r *http.Request, pathToFile, fileName string) error {
	fp := path.Join(pathToFile, fileName)
	fileToServe := filepath.Clean(fp)
	r.Header.Set("Content-Type", fmt.Sprintf("attachment; file=\"%s\"", fileName))
	http.ServeFile(w, r, fileToServe)
	return nil
}

func (g *Gojango) Error404(w http.ResponseWriter, r *http.Request) {
	g.ErrorStatus(w, http.StatusNotFound)
}

func (g *Gojango) Error500(w http.ResponseWriter, r *http.Request) {
	g.ErrorStatus(w, http.StatusInternalServerError)
}

func (g *Gojango) Error401(w http.ResponseWriter, r *http.Request) {
	g.ErrorStatus(w, http.StatusUnauthorized)
}

func (g *Gojango) Error403(w http.ResponseWriter, r *http.Request) {
	g.ErrorStatus(w, http.StatusForbidden)
}

func (g *Gojango) ErrorStatus(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
