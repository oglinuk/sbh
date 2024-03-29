package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
	"strings"

	"github.com/oglinuk/sbh"
)

var (
	algorithm = flag.String("a", "", "Hashing algorithm to use, defaults to sha256")
	uptimes   = flag.Int("ut", 0, "Number of letters to make uppercase")
	symbols   = flag.String("s", "", "Symbols to add to SBH")
	length    = flag.Int("l", 0, "Length of the returned string")
)

func doAgain() bool {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nGenerate another? [y/N]: ")
	scanner.Scan()
	again := strings.ToLower(scanner.Text())

	return again == "y"
}

func genSBH() {
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
		Length:         *length,
	}

	if *algorithm == "" {
		fmt.Println("No algorithm specified with -a, defaulting to sha256 ...")
	}

	sTime := time.Now()

	fmt.Printf("SBH: %s\nElapsed time: %v\n",
		sbh.Generate(secbaehash), time.Since(sTime))

	scanner.Scan()
}

func main() {
	flag.Parse()

	cont := true
	for {
		if !cont {
			break
		}
		genSBH()
		cont = doAgain()
	}
}

