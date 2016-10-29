package main

import (
	"github.com/MarcosSegovia/MyWins/src/wins/domain"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler apiHandlerFunc
}

type Routes []Route

func NewRouter(mongoDB *mgo.Session, dbInfo *domain.DBInfo) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, singleRoute := range routes {

		router.Methods(singleRoute.Method).Path(singleRoute.Path).Name(singleRoute.Name).Handler(newGeneralHandler(mongoDB, dbInfo, singleRoute.Handler))
	}

	return router
}

var routes = Routes{
	Route{"Get Wins", "GET", "/wins", GetAllWins},
	Route{"Get Wins", "GET", "/fails", GetAllFails},
	Route{"Post Win", "POST", "/wins/add", AddWin},
	Route{"Post Win", "POST", "/fails/add", AddFail},
	Route{"Hi", "GET", "/", Welcome},
}
