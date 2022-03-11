package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	//"github.com/gin-gonic/contrib/cors"
	//"github.com/gin-gonic/gin"
)

var (
	tpl *template.Template
)

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9001"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/sbh", sbhHandler)

	httpFS := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", httpFS))

	log.Printf("Server is running at 0.0.0.0:%s ...", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
