package mongo

import (
	"log"

	"github.com/MarcosSegovia/MyWins/src/wins/domain"
)

type MongoApiClient struct {
	dbClient *MongoClient
}

func NewMongoApiClient() MongoApiClient{
	return MongoApiClient{NewMongoClient()}
}

func (client MongoApiClient) GetWins() ([]*domain.Entry, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	collection := session.DB(client.dbClient.DbInfo.DB).C(client.dbClient.DbInfo.WinColl)
	var wins []*domain.Entry
	err := collection.Find(nil).All(&wins)
	if err != nil {
		log.Println(err)
	}
	return wins, err
}

func (client MongoApiClient) GetFails() ([]*domain.Entry, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	collection := session.DB(client.dbClient.DbInfo.DB).C(client.dbClient.DbInfo.FailColl)
	var fails []*domain.Entry
	err := collection.Find(nil).All(&fails)
	if err != nil {
		log.Println(err)
	}
	return fails, err
}

func (client MongoApiClient) AddWin() error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	collection := session.DB(client.dbClient.DbInfo.DB).C(client.dbClient.DbInfo.WinColl)
	err := collection.Insert(domain.NewEntry())
	if err != nil {
		log.Println(err)
	}
	return err
}

func (client MongoApiClient) AddFail() error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	collection := session.DB(client.dbClient.DbInfo.DB).C(client.dbClient.DbInfo.FailColl)
	err := collection.Insert(domain.NewEntry())
	if err != nil {
		log.Println(err)
	}
	return err
}
