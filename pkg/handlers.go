package pkg

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const formName = "json_input"

type ThreadService struct {
	mainHtmlPath string
	tmpl         *template.Template
}

func NewThreadService(mainHtmlPath string, tmpl *template.Template) *ThreadService {
	return &ThreadService{
		mainHtmlPath: mainHtmlPath,
		tmpl:         tmpl,
	}
}

func (s *ThreadService) HomePage() http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "%s", s.mainHtmlPath)
	})
}

func (s *ThreadService) Render() http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing form: %v", err), http.StatusInternalServerError)
			return
		}

		formValue := r.PostFormValue(formName)

		data := Thread{}
		if err := json.Unmarshal([]byte(formValue), &data); err != nil {
			http.Error(w, fmt.Sprintf("error parsing request body: %v", err), http.StatusBadRequest)
			return
		}

		templateData := &HtmlTemplate{}
		templateData.FromAPI(data)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		s.tmpl.Execute(w, templateData)
	})
}
