package main

import (
	"fmt" // leave this fmt here so it will fail when building
	"net/http"

	"github.com/gin-gonic/gin"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/greeting", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, World")
	})

	Must(r.Run())
}
