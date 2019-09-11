package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/OGLinuk/go-sbh/sbh"
)

var (
	plainText = flag.String("pt", "", "Plaintext to encode")
	nRots     = flag.Int64("nr", 0, "Number of rotations to apply")
	seed      = flag.Int64("s", 0, "Seed for random int generation")
	web       = flag.Bool("w", false, "Serve SBH over http")
)

func main() {
	flag.Parse()

	print("\033[H\033[2J") //clear terminal

	if *plainText != "" && *nRots != 0 && *seed != 0 {
		sTime := time.Now()
		fmt.Printf("SBH: %s\nElapsed time: %v\n",
			sbh.Generate(*plainText, *nRots, *seed), time.Since(sTime))
	} else if *web {
		http.HandleFunc("/sbh", serveSBH)
		fmt.Println("Starting server on 9001 ...")
		http.ListenAndServe("0.0.0.0:9001", nil)
	} else {
		fmt.Println("Please provide a flag. For more information use -h")
	}
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
