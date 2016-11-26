package main

import (
        "net/http"

        "github.com/gorilla/mux"
)

// Route represents a route url with a handler associated in both authentication and client server.
type Route struct {
        Name    string
        Method  string
        Path    string
        Handler http.HandlerFunc
}

// Routes are the server group of routes to be able to access.
type Routes []Route

var serverRoutes = Routes{
        Route{"Login into MyWins", "GET", "/authorize", Authorization},
        Route{"Get Access token", "POST", "/token", AccessToken},
        Route{"Get Wins", "GET", "/wins", GetAllWins},
        Route{"Get Wins", "GET", "/fails", GetAllFails},
        Route{"Post Win", "POST", "/wins/add", AddWin},
        Route{"Post Win", "POST", "/fails/add", AddFail},
        Route{"Hi", "GET", "/", Welcome},
        Route{"Information Endpoint", "GET", "/info", Information},
}

var clientRoutes = Routes{
        Route{"Get Login into MyWins", "GET", "/login", Login},
        Route{"Post Login into MyWins", "POST", "/login", LoginPost},
        Route{"Get Access token", "GET", "/accesstoken", AuthForAccessToken},
}

// NewServerRouter instantiates a new Authentication and MyWins Router.
func NewServerRouter() *mux.Router {
        router := mux.NewRouter().StrictSlash(true)

        for _, singleRoute := range serverRoutes {

                router.Methods(singleRoute.Method).Path(singleRoute.Path).Name(singleRoute.Name).Handler(singleRoute.Handler)
        }

        return router
}

// NewClientRouter instantiates a new Client Router
func NewClientRouter() *mux.Router {
        router := mux.NewRouter().StrictSlash(true)

        for _, singleRoute := range clientRoutes {

                router.Methods(singleRoute.Method).Path(singleRoute.Path).Name(singleRoute.Name).Handler(singleRoute.Handler)
        }

        return router
}
