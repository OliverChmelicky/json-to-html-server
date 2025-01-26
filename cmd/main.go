package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/OliverChmelicky/json-to-html-server/pkg"
)

func main() {
	var port int
	var templatePath string
	var mainHtmlPath string
	flag.IntVar(&port, "p", 8080, "Port on which the server runs.")
	flag.StringVar(&templatePath, "t", "./templates/threat.html.tmpl", "Path to template file for generating page for threads.")
	flag.StringVar(&mainHtmlPath, "m", "./templates/main.html", "Path to main path handling root URL path.")
	flag.Parse()

	// Parse and execute the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("couldd not parse temlate: %s", err)
		return
	}

	mainPage, err := os.ReadFile(mainHtmlPath)
	if err != nil {
		log.Fatalf("error opening file: %s", err)
		return
	}

	threadService := pkg.NewThreadService(tmpl, string(mainPage))

	http.HandleFunc("GET /", threadService.HomePage())
	http.HandleFunc("POST /render", threadService.LoggerMiddleware(threadService.Render()))

	portStr := strconv.Itoa(port)
	log.Println("Server listens on port: ", port)
	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
