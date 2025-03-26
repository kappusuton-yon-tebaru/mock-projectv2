package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

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
	r := gin.Default()

	delayedServiceUrl, err := url.Parse(Getenv("DELAYED_SERVICE_URL", ""))
	Must(err)
	delayedServiceRp := httputil.NewSingleHostReverseProxy(delayedServiceUrl)

	r.Any("/delayed-service/*any", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/delayed-service")
		delayedServiceRp.ServeHTTP(w, r)
	}))

	service1ServiceUrl, err := url.Parse(Getenv("SERVICE1_SERVICE_URL", ""))
	Must(err)
	service1ServiceRp := httputil.NewSingleHostReverseProxy(service1ServiceUrl)

	r.Any("/service1/*any", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/service1")
		service1ServiceRp.ServeHTTP(w, r)
	}))

	Must(r.Run())
}
