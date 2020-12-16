package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/OGLinuk/sbh/sbh"
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

	hash := sbh.Generate(secbaehash)

	fmt.Printf("SBH: %s\nElapsed time: %v\n",
		hash, time.Since(sTime))
}

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
