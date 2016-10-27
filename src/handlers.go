package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcosSegovia/MyWins/src/wins/domain"
	"gopkg.in/mgo.v2"
)

type apiHandlerFunc func(http.ResponseWriter, *http.Request, *domain.MyWinsAPI)

//GeneralHandler is used to wrap all other handlers
type GeneralHandler struct {
	mongoDB     *mgo.Session
	dbInfo      *domain.DBInfo
	handlerFunc apiHandlerFunc
}

func newGeneralHandler(mongoDB *mgo.Session, dbInfo *domain.DBInfo, handler apiHandlerFunc) *GeneralHandler {
	return &GeneralHandler{
		mongoDB:     mongoDB,
		dbInfo:      dbInfo,
		handlerFunc: handler,
	}
}

func (gh *GeneralHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := gh.mongoDB.Copy()
	defer session.Close()
	api := domain.NewApi(session, gh.dbInfo)
	gh.handlerFunc(w, r, api)
}

func Welcome(w http.ResponseWriter, r *http.Request, api *domain.MyWinsAPI) {
	fmt.Fprintln(w, "Welcome!!!")
}

func GetAllWins(w http.ResponseWriter, r *http.Request, api *domain.MyWinsAPI) {
	wins, err := api.FindAllWins()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(wins))
}

func GetAllFails(w http.ResponseWriter, r *http.Request, api *domain.MyWinsAPI) {
	fails, err := api.FindAllFails()
	if err != nil {
		http.Error(w, domain.GeneralError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(fails))
}

func AddWin(w http.ResponseWriter, r *http.Request, api *domain.MyWinsAPI) {
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
func AddFail(w http.ResponseWriter, r *http.Request, api *domain.MyWinsAPI) {
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
