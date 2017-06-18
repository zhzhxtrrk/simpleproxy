package main

import (
	"github.com/elazarl/goproxy"
	"net/http"
	"os"
	"log"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	portEnv, portEnvPresent := os.LookupEnv("PORT")
	if !portEnvPresent {
		portEnv = "8000"
	}
	endpoint := ":" + portEnv

	log.Println("Starting proxy server on", endpoint)
	log.Fatal(http.ListenAndServe(endpoint, proxy))
}
