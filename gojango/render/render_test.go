package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct{
	name string
	renderer string
	template string
	errorExpected bool
	errorMessage string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-existing jet template"},
	{"jet_page", "jet", "home", false, "error rendering go template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-existing jet template"},
	{"invalid_render_engine", "foo", "no-file", true, "no error rendering invalid template engine"},
}

func TestRender_Page(t *testing.T) {
	for _, e := range pageData{
		r, err := http.NewRequest("GET", "/test-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

