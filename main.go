package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./routes"
)

const LISTEN_ADDRESS_ENV = "LISTEN"
const LISTEN_PORT_ENV = "PORT"
const DEFAULT_LISTEN_ADDRESS = "0.0.0.0"
const DEFAULT_LISTEN_PORT = "80"

func main() {
	httpListenAddress := getEnv(LISTEN_ADDRESS_ENV, DEFAULT_LISTEN_ADDRESS)
	httpListenPort := getEnv(LISTEN_PORT_ENV, DEFAULT_LISTEN_PORT)

	srv := &http.Server{
		Handler:     routes.Router(),
		Addr:        httpListenAddress + ":" + httpListenPort,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getEnv(name, defaultVal string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}
