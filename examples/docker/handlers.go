package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oglinuk/sbh"
)

type resp struct {
	SBH      string
	TimeTook time.Duration
}

func restHandler(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		var secbaehash sbh.SBH

		err := ctx.BindJSON(&secbaehash)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("ctx.BindJSON: %s", err.Error()),
			})
			return
		}

		if secbaehash.Algorithm == "" {
			secbaehash.Algorithm = "sha256"
		}

		sTime := time.Now()

		ctx.JSON(http.StatusOK, gin.H{
			"sbh":       sbh.Generate(secbaehash),
			"time-took": time.Since(sTime),
		})
	}
}

func uiHandler(ctx *gin.Context) {
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

		l := ctx.PostForm("length")
		if l == "" {
			l = "0"
		}

		length, err := strconv.Atoi(l)
		if err != nil {
			http.Error(ctx.Writer, fmt.Sprintf("length: %s", err.Error()), 500)
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
