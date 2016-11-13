package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

var serverRoutes = Routes{
	Route{"Login into MyWins", "GET", "/authorize", Authorization},
	Route{"Get Access token", "POST", "/token", AccessToken},
	Route{"Get Wins", "GET", "/wins", GetAllWins},
	Route{"Get Wins", "GET", "/fails", GetAllFails},
	Route{"Post Win", "POST", "/wins/add", AddWin},
	Route{"Post Win", "POST", "/fails/add", AddFail},
	Route{"Hi", "GET", "/", Welcome},
}

var clientRoutes = Routes{
	Route{"Login into MyWins", "GET", "/", Login},
	Route{"Get Access token", "GET", "/accesstoken", AuthForAccessToken},
}

func NewServerRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, singleRoute := range serverRoutes {

		router.Methods(singleRoute.Method).Path(singleRoute.Path).Name(singleRoute.Name).Handler(singleRoute.Handler)
	}

	return router
}

func NewClientRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, singleRoute := range clientRoutes {

		router.Methods(singleRoute.Method).Path(singleRoute.Path).Name(singleRoute.Name).Handler(singleRoute.Handler)
	}

	return router
}
