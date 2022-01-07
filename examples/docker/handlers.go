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

func restHandler(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		var secbaehash sbh.SBH

		err := ctx.BindJSON(&secbaehash)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": fmt.Sprintf("ctx.BindJSON: %s", err.Error()),
			})
			return
		}

		if secbaehash.Algorithm == "" {
			secbaehash.Algorithm = "sha256"
		}

		if secbaehash.Uppercase && secbaehash.UppercaseTimes == 0 {
			secbaehash.UppercaseTimes = 1
		}

		sTime := time.Now()

		ctx.JSON(http.StatusOK, gin.H{
			"sbh": sbh.Generate(secbaehash),
			"time-complexity": time.Since(sTime),
		})
	}
}

func uiHandler(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		plaintext := ctx.PostForm("plaintext")
		nrots, err := strconv.ParseInt(ctx.PostForm("nrots"), 10, 64)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), 500)
		}

		seed, err := strconv.ParseInt(ctx.PostForm("seed"), 10, 64)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), 500)
		}

		algorithm := ctx.PostForm("algorithm")

		uppercase := false
		var uptimes int64
		uptimes = 0
		if ctx.PostForm("uppercase") != "" {
			uppercase = true
			uptimes, err = strconv.ParseInt(ctx.PostForm("uptimes"), 10, 64)
			if err != nil {
				http.Error(ctx.Writer, err.Error(), 500)
			}
		}

		symbols := ctx.PostForm("symbols")

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
