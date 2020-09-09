package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"net/http"
	"strconv"
	"time"

	"github.com/OGLinuk/go-sbh/sbh"
)

var (
	web = flag.Bool("w", false, "Serve SBH over http")
)

func SBH() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Plaintext: ")
	scanner.Scan()
	plainText := scanner.Text()

	var nRots int64
	fmt.Printf("Number of Rotations: ")
	fmt.Scan(&nRots)

	var seed int64
	fmt.Printf("Seed: ")
	fmt.Scan(&seed)
	
	print("\033[H\033[2J")

	sTime := time.Now()
	fmt.Printf("SBH: %s\nElapsed time: %v\n",
		sbh.Generate(plainText, nRots, seed), time.Since(sTime))
}

// <IP Address>:9001/sbh?plaintext=test&nrots=1729&seed=42
func serveSBH(w http.ResponseWriter, r *http.Request) {
	sTime := time.Now()
	pt := r.FormValue("plaintext")
	nr, err := strconv.ParseInt(r.FormValue("nrots"), 10, 64)
	if err != nil {
		return
	}

	s, err := strconv.ParseInt(r.FormValue("seed"), 10, 64)
	if err != nil {
		return
	}

	fmt.Fprintln(w, fmt.Sprintf("SBH: %s\nElapsed time: %v",
		sbh.Generate(pt, nr, s), time.Since(sTime)))
}

func main() {
	flag.Parse()

	print("\033[H\033[2J") //clear terminal

	if *web {
		http.HandleFunc("/sbh", serveSBH)
		fmt.Println("Starting server on 9001 ...")
		http.ListenAndServe("0.0.0.0:9001", nil)
	} else {
		SBH()
	}
}

