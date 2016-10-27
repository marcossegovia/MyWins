package domain

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
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

type DBInfo struct {
	DB       string
	WinColl  string
	FailColl string
}

func NewDBInfo(db, winColl, failColl string) *DBInfo {
	return &DBInfo{
		DB:       db,
		WinColl:  winColl,
		FailColl: failColl,
	}
}

type MyWinsAPI struct {
	session *mgo.Session
	dbInfo  *DBInfo
}

func NewApi(session *mgo.Session, dbInfo *DBInfo) *MyWinsAPI {
	api := new(MyWinsAPI)
	api.session = session
	api.dbInfo = dbInfo
	return api
}

func (api *MyWinsAPI) FindAllWins() ([]*Entry, error) {
	collection := api.session.DB(api.dbInfo.DB).C(api.dbInfo.WinColl)
	var wins []*Entry
	err := collection.Find(nil).All(&wins)
	if err != nil {
		log.Println(err)
	}
	return wins, err
}

func (api *MyWinsAPI) FindAllFails() ([]*Entry, error) {
	collection := api.session.DB(api.dbInfo.DB).C(api.dbInfo.FailColl)
	var fails []*Entry
	err := collection.Find(nil).All(&fails)
	if err != nil {
		log.Println(err)
	}
	return fails, err
}

func (api *MyWinsAPI) AddWin() error {
	collection := api.session.DB(api.dbInfo.DB).C(api.dbInfo.WinColl)
	err := collection.Insert(NewEntry())
	if err != nil {
		log.Println(err)
	}
	return err
}

func (api *MyWinsAPI) AddFail() error {
	collection := api.session.DB(api.dbInfo.DB).C(api.dbInfo.FailColl)
	err := collection.Insert(NewEntry())
	if err != nil {
		log.Println(err)
	}
	return err
}
