package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"./wins/domain"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!!!")
}

func GetAllWins(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi()
	wins_response, err := api.FindAllWins()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(wins_response.Success))
}

func GetAllFails(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi()
	wins_response, err := api.FindAllWins()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(wins_response.Fails))
}

func AddWin(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi()
	err := api.AddWin()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response_string := []string{"Win added =), Keep it up."};
	w.Write(buildResponse(response_string))
}
func AddFail(w http.ResponseWriter, r *http.Request) {
	api := domain.NewApi()
	err := api.AddFail()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response_string := []string{"Fail added =(, Sorry to hear that."};
	w.Write(buildResponse(response_string))
}

func buildResponse(slice_of_times []string) []byte {

	json_encoded_times, err := json.Marshal(slice_of_times)
	if err != nil {

	}
	return json_encoded_times
}
