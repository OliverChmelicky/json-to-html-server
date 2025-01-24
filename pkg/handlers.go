package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Url: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func RegisterHomePage(mainHtmlPath string) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {
		html, err := os.ReadFile(mainHtmlPath)
		if err != nil {
			log.Println("Error opening file:", err)
			http.Error(w, fmt.Sprintf("Error opening file %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, string(html))
	})
}

func BarHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Bar Handler")
		fmt.Fprintf(w, "Bar Handler")
	})
}
