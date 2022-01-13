package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
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

	grouter := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	grouter.GET("/", indexHandler)
	grouter.POST("/", sbhHandler)
	grouter.GET("/sbh", sbhHandler)
	grouter.Static("/static/", "./static")

	log.Printf("Server is running at 0.0.0.0:%s ...", PORT)
	log.Fatal(grouter.Run(fmt.Sprintf("0.0.0.0:%s", PORT)), nil)
}
