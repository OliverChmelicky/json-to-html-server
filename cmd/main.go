package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/OliverChmelicky/json-to-html-server/pkg"
)

const mainHtml = "./templates/main.html" // TODO extract as env var

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "specify port on which the server runs")
	flag.Parse()

	mux := http.NewServeMux()

	homePage := pkg.RegisterHomePage(mainHtml)
	homePageHandler := http.HandlerFunc(homePage)
	mux.Handle("/", pkg.LoggerMiddleware(homePageHandler))

	mux.Handle("/bar", pkg.LoggerMiddleware(pkg.BarHandler()))

	portStr := strconv.Itoa(port)
	log.Println("Server listens on port: ", port)
	log.Fatal(http.ListenAndServe(":"+portStr, mux))
}
