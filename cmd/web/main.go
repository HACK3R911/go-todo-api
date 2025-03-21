package main

import (
	"log"
	"net/http"
)

func main() {
	srv := new(http.Server)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
