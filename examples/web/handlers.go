package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/oglinuk/sbh"
	"github.com/gin-gonic/gin"
)

type resp struct {
	SBH            string
	TimeComplexity time.Duration
}

func sbhHandler(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		plaintext := ctx.PostForm("plaintext")
		nrots, err := strconv.ParseInt(ctx.PostForm("nrots"), 10, 64)
		if err != nil {
			http.Error(ctx.Writer, fmt.Sprintf("nrots: %s", err.Error()), 500)
		}

		seed, err := strconv.ParseInt(ctx.PostForm("seed"), 10, 64)
		if err != nil {
			http.Error(ctx.Writer, fmt.Sprintf("seed: %s", err.Error()), 500)
		}

		algorithm := ctx.PostForm("algorithm")
		if algorithm == "" {
			algorithm = "sha256"
		}

		ut := ctx.PostForm("uppercasetimes")
		if ut == "" {
			ut = "0"
		}

		uptimes, err := strconv.Atoi(ut)
		if err != nil {
			http.Error(ctx.Writer, fmt.Sprintf("uppercasetimes: %s", err.Error()), 500)
		}

		symbols := ctx.PostForm("symbols")

		secbaehash := sbh.SBH{
			Plaintext:      plaintext,
			NRots:          nrots,
			Seed:           seed,
			Algorithm:      algorithm,
			UppercaseTimes: uptimes,
			Symbols:        symbols,
		}

		sTime := time.Now()

		if err := tpl.ExecuteTemplate(
			ctx.Writer,
			"index.html",
			resp{sbh.Generate(secbaehash), time.Since(sTime)},
		); err != nil {
			http.Error(ctx.Writer, err.Error(), 500)
		}
	}
}

func indexHandler(ctx *gin.Context) {
	tpl = template.Must(template.ParseGlob("templates/*"))
	err := tpl.ExecuteTemplate(ctx.Writer, "index.html", nil)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}
