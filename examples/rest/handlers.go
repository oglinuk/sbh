package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oglinuk/sbh"
)

func sbhHandler(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		var secbaehash sbh.SBH

		err := ctx.BindJSON(&secbaehash)
		if err != nil {
			log.Println(err.Error())
		}

		if secbaehash.Algorithm == "" {
			secbaehash.Algorithm = "sha256"
		}

		sTime := time.Now()

		ctx.JSON(http.StatusOK, gin.H{
			"sbh":             sbh.Generate(secbaehash),
			"time-complexity": fmt.Sprintf("%s", time.Since(sTime)),
		})
	}
}
