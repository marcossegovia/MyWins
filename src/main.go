package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	e := http.ListenAndServe(":8080", router)

	if e != nil {

		log.Fatal(e)
	}
}
