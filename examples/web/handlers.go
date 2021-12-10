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

var (
	tpl *template.Template
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

		tpl = template.Must(template.ParseGlob("templates/*"))
		if err := tpl.ExecuteTemplate(
			ctx.Writer,
			"index.html",
			resp{sbh.Generate(secbaehash), time.Since(sTime)},
		); err != nil {
			http.Error(ctx.Writer, err.Error(), 500)
		}
	} else {
		plaintext := ctx.Query("plaintext")

		nrots, err := strconv.ParseInt(ctx.Query("nrots"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to strconv.ParseInt(numrots): %s", err.Error()),
			})
		}

		seed, err := strconv.ParseInt(ctx.Query("seed"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to strconv.ParseInt(seed): %s", err.Error()),
			})
		}

		algorithm := ctx.Query("algorithm")

		uppercase, err := strconv.ParseBool(ctx.Query("uppercase"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to strconv.ParseBool(uppercase): %s", err.Error()),
			})
		}

		uptimes, err := strconv.ParseInt(ctx.Query("uptimes"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to strconv.ParseInt(uptimes): %s", err.Error()),
			})
		}

		symbols := ctx.Query("symbols")

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

		ctx.JSON(http.StatusOK, gin.H{
			"sbh":             sbh.Generate(secbaehash),
			"time-complexity": fmt.Sprintf("%s", time.Since(sTime)),
		})
	}
}

func indexHandler(ctx *gin.Context) {
	tpl = template.Must(template.ParseGlob("templates/*"))
	err := tpl.ExecuteTemplate(ctx.Writer, "index.html", nil)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}
