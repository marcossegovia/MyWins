package main

import (
	"log"
	"net/http"
)

func main() {

	BootstrapClient()

	serverRouter := NewServerRouter()
	clientRouter := NewClientRouter()

	go http.ListenAndServe(":8081", clientRouter)
	e := http.ListenAndServe(":8080", serverRouter)

	if e != nil {

		log.Fatal(e)
	}
}
