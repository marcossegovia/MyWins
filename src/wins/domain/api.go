package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Entry struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Time int64         `json:"time" bson:"time"`
}

func NewEntry() *Entry {
	entry := new(Entry)
	entry.ID = bson.NewObjectId()
	entry.Time = time.Now().Unix()
	return entry
}

var api *MyWinsAPI

type MyWinsAPI struct {
	dbClient PersistenceApiClient
}

func NewApi(client PersistenceApiClient) *MyWinsAPI {

	if api != nil {
		return api
	}
	api := &MyWinsAPI{client}
	return api
}

func (api *MyWinsAPI) FindAllWins() ([]*Entry, error) {
	return api.dbClient.GetWins()
}

func (api *MyWinsAPI) FindAllFails() ([]*Entry, error) {
	return api.dbClient.GetFails()
}

func (api *MyWinsAPI) AddWin() error {
	return api.dbClient.AddWin()
}

func (api *MyWinsAPI) AddFail() error {
	return api.dbClient.AddFail()
}
