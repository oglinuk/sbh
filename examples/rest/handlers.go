package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/oglinuk/sbh"
	"github.com/gin-gonic/gin"
)

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

		ctx.JSON(http.StatusOK, gin.H{
			"sbh": sbh.Generate(secbaehash),
			"time-complexity": time.Since(sTime),
		})
	}
}