package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcosSegovia/MyWins/src/wins/domain"
	"github.com/MarcosSegovia/MyWins/src/wins/infrastructure/mongo"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!!!")
}

func GetAllWins(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi(mongo.NewMongoApiClient())
	wins, err := api.FindAllWins()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(wins))
}

func GetAllFails(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi(mongo.NewMongoApiClient())
	fails, err := api.FindAllFails()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(fails))
}

func AddWin(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi(mongo.NewMongoApiClient())
	err := api.AddWin()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response_string := []string{"Win added =), Keep it up."}
	w.Write(buildResponse(response_string))
}
func AddFail(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi(mongo.NewMongoApiClient())
	err := api.AddFail()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response_string := []string{"Fail added =(, Sorry to hear that."}
	w.Write(buildResponse(response_string))
}

func buildResponse(response interface{}) []byte {

	resp, err := json.Marshal(response)
	if err != nil {

	}
	return resp
}
