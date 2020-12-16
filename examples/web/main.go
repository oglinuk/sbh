package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"git.sr.ht/~oglinuk/sbh"
)

func serveSBH(w http.ResponseWriter, r *http.Request) {
	algorithm := r.FormValue("algorithm")
	plaintext := r.FormValue("plaintext")
	symbols := r.FormValue("symbols")

	nrots, err := strconv.ParseInt(r.FormValue("nrots"), 10, 64)
	if err != nil {
		return
	}

	seed, err := strconv.ParseInt(r.FormValue("seed"), 10, 64)
	if err != nil {
		return
	}

	uppercase, err := strconv.ParseBool(r.FormValue("uppercase"))
	if err != nil {
		return
	}

	uptimes, err := strconv.ParseInt(r.FormValue("uptimes"), 10, 64)
	if err != nil {
		return
	}

	secbaehash := sbh.SBH{
		Plaintext:      plaintext,
		NRots:          nrots,
		Seed:           seed,
		Algorithm:      algorithm,
		Uppercase:      uppercase,
		UppercaseTimes: int(uptimes),
		Symbols:        symbols,
	}

	sTime := time.Now()

	hash := sbh.Generate(secbaehash)

	if algorithm == "" {
		fmt.Fprintln(w, fmt.Sprintf("No algorithm specified with -a, defaulting to sha256 ...\nSBH: %s\nElapsed time: %v",
			hash, time.Since(sTime)))
	} else {
		fmt.Fprintln(w, fmt.Sprintf("SBH: %s\nElapsed time: %v", hash, time.Since(sTime)))
	}
}

func main() {
	http.HandleFunc("/sbh", serveSBH)
	fmt.Println("Starting server on 9001 ...")
	http.ListenAndServe("0.0.0.0:9001", nil)
}
