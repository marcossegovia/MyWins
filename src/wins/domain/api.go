package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Entry is the main class to encapsulate a single win or fail.
type Entry struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Time int64         `json:"time" bson:"time"`
}

// NewEntry creates a new Entry with the current time.
func NewEntry() *Entry {
	entry := new(Entry)
	entry.ID = bson.NewObjectId()
	entry.Time = time.Now().Unix()
	return entry
}

var api *MyWinsAPI

// MyWinsAPI is the class that implements all MyWins endpoints.
type MyWinsAPI struct {
	dbClient PersistenceAPIClient
}

// NewAPI instantiates MyWins API.
func NewAPI(client PersistenceAPIClient) *MyWinsAPI {
	if api != nil {
		return api
	}
	api := &MyWinsAPI{client}
	return api
}

// FindAllWins retrieves all wins.
func (api *MyWinsAPI) FindAllWins() ([]*Entry, error) {
	return api.dbClient.GetWins()
}

// FindAllFails retrieves all fails.
func (api *MyWinsAPI) FindAllFails() ([]*Entry, error) {
	return api.dbClient.GetFails()
}

// AddWin submits a win for the current day.
func (api *MyWinsAPI) AddWin() error {
	return api.dbClient.AddWin()
}

// AddFail submits a fail for the current day.
func (api *MyWinsAPI) AddFail() error {
	return api.dbClient.AddFail()
}
