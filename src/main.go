package main

import (
	"fmt"
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

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
