package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, singleRoute := range routes {

		router.Methods(singleRoute.Method).Path(singleRoute.Path).Name(singleRoute.Name).Handler(singleRoute.Handler)
	}

	return router
}

var routes = Routes{
	Route{"Get Wins", "GET", "/wins", Welcome},
	Route{"Post Win", "POST", "/wins/add", Welcome},
	Route{"Hi", "GET", "/", Welcome},
}
