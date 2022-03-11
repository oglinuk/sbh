package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/oglinuk/sbh"
)

type resp struct {
	SBH      string
	TimeTook time.Duration
}

func sbhHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		plaintext := r.FormValue("plaintext")
		nrots, err := strconv.ParseInt(r.FormValue("nrots"), 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("nrots: %s", err.Error()), 500)
		}

		seed, err := strconv.ParseInt(r.FormValue("seed"), 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("seed: %s", err.Error()), 500)
		}

		algorithm := r.FormValue("algorithm")
		if algorithm == "" {
			algorithm = "sha256"
		}

		ut := r.FormValue("uppercasetimes")
		if ut == "" {
			ut = "0"
		}

		uptimes, err := strconv.Atoi(ut)
		if err != nil {
			http.Error(w, fmt.Sprintf("uppercasetimes: %s", err.Error()), 500)
		}

		symbols := r.FormValue("symbols")

		l := r.FormValue("length")
		if l == "" {
			l = "0"
		}

		length, err := strconv.Atoi(l)
		if err != nil {
			http.Error(w, fmt.Sprintf("length: %s", err.Error()), 500)
		}

		secbaehash := sbh.SBH{
			Plaintext:      plaintext,
			NRots:          nrots,
			Seed:           seed,
			Algorithm:      algorithm,
			UppercaseTimes: uptimes,
			Symbols:        symbols,
			Length:         length,
		}

		sTime := time.Now()

		if err := tpl.ExecuteTemplate(
			w,
			"index.html",
			resp{sbh.Generate(secbaehash), time.Since(sTime)},
		); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseGlob("templates/*"))
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
