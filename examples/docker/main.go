package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"git.sr.ht/~oglinuk/sbh"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	algorithm = flag.String("a", "", "Hashing algorithm to use, defaults to sha256")
	uppercase = flag.Bool("u", false, "Make a letter of SBH uppercase")
	uptimes   = flag.Int("ut", 1, "Number of letters to make uppercase")
	symbols   = flag.String("s", "", "Symbols to add to SBH")
	web       = flag.Bool("w", false, "Serve SBH over http")
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
		Uppercase:      *uppercase,
		UppercaseTimes: *uptimes,
		Symbols:        *symbols,
	}

	print("\033[H\033[2J")
	if *algorithm == "" {
		fmt.Println("No algorithm specified with -a, defaulting to sha256 ...")
	}

	sTime := time.Now()

	fmt.Printf("SBH: %s\nElapsed time: %v\n",
		sbh.Generate(secbaehash), time.Since(sTime))
}

func main() {
	flag.Parse()

	print("\033[H\033[2J") //clear terminal

	if *web {
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
	} else {
		SBH()
	}
}
