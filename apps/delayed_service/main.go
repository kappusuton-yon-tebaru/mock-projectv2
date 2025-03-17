package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Getenv(key, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}

func main() {
	startUpdelay, err := strconv.Atoi(Getenv("START_UP_DELAY", "5"))
	Must(err)

	time.Sleep(time.Duration(startUpdelay) * time.Second)

	r := gin.Default()

	r.GET("/hc", func(ctx *gin.Context) {
		ctx.Status(200)
	})

	Must(r.Run())
}
