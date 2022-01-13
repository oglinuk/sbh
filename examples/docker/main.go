package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/oglinuk/sbh"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	algorithm = flag.String("a", "", "Hashing algorithm to use, defaults to sha256")
	uptimes   = flag.Int("ut", 0, "Number of letters to make uppercase")
	symbols   = flag.String("s", "", "Symbols to add to SBH")
	web       = flag.Bool("w", false, "Serve SBH over http")

	tpl *template.Template
)

func SBH() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Plaintext: ")
	scanner.Scan()
	plaintext := scanner.Text()

	var nrots int64
	fmt.Printf("Number of Rotations: ")
	fmt.Scan(&nrots)

	var seed int64
	fmt.Printf("Seed: ")
	fmt.Scan(&seed)

	secbaehash := sbh.SBH{
		Plaintext:      plaintext,
		NRots:          nrots,
		Seed:           seed,
		Algorithm:      *algorithm,
		UppercaseTimes: *uptimes,
		Symbols:        *symbols,
	}

	if *algorithm == "" {
		fmt.Println("No algorithm specified with -a, defaulting to sha256 ...")
	}

	sTime := time.Now()

	fmt.Printf("SBH: %s\nElapsed time: %v\n",
		sbh.Generate(secbaehash), time.Since(sTime))
}

func main() {
	flag.Parse()

	if *web {
		PORT := os.Getenv("PORT")
		if PORT == "" {
			PORT = "9001"
		}

		tpl = template.Must(template.ParseGlob("templates/*"))

		grouter := gin.Default()

		corsConfig := cors.DefaultConfig()
		corsConfig.AllowAllOrigins = true

		grouter.GET("/", indexHandler)
		grouter.POST("/", uiHandler)
		grouter.POST("/sbh", restHandler)
		grouter.Static("/static/", "./static")

		log.Printf("Server is running at :%s ...", PORT)
		log.Fatal(grouter.Run(fmt.Sprintf(":%s", PORT)), nil)
	} else {
		SBH()
	}
}
