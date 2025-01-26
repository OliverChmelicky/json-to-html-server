package pkg

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ThreadService struct {
	tmpl         *template.Template
	mainHtmlPath string
}

func NewThreadService(tmpl *template.Template, mainHtmlPath string) *ThreadService {
	return &ThreadService{
		tmpl:         tmpl,
		mainHtmlPath: mainHtmlPath,
	}
}

func (s *ThreadService) LoggerMiddleware(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		log.Println("Url: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (s *ThreadService) HomePage() http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, s.mainHtmlPath)
	})
}

func (s *ThreadService) Render() http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("error parsing form: %s", err), http.StatusInternalServerError)
			return
		}

		formValue := r.PostFormValue("json_input")

		data := Thread{}
		if err := json.Unmarshal([]byte(formValue), &data); err != nil {
			http.Error(w, fmt.Sprintf("error parsing request body: %s", err), http.StatusBadRequest)
			return
		}

		templateData := &HtmlTemplate{}
		templateData.FromAPI(data)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		s.tmpl.Execute(w, templateData)
	})
}
