package mongo

import (
	"log"

	"github.com/marcossegovia/MyWins/src/wins/domain"
)

// MongoAPIClient is the instance that implements MyWins PersistenceAPIClient.
type MongoAPIClient struct {
	dbClient *MongoClient
}

// NewMongoAPIClient instantiates a Mongo struct that implements all MyWins PersistenceAPIClient and holds an instance of a Mongodb Connection.
func NewMongoAPIClient() MongoAPIClient {
	return MongoAPIClient{NewMongoClient()}
}

// GetWins retrieves all wins from the mongodb wins collection.
func (client MongoAPIClient) GetWins() ([]*domain.Entry, error) {
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

// GetFails retrieves all fails from the mongodb fails collection.
func (client MongoAPIClient) GetFails() ([]*domain.Entry, error) {
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

// AddWin persist a new Entry into mongodb wins collection.
func (client MongoAPIClient) AddWin() error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	collection := session.DB(client.dbClient.DbInfo.DB).C(client.dbClient.DbInfo.WinColl)
	err := collection.Insert(domain.NewEntry())
	if err != nil {
		log.Println(err)
	}
	return err
}

// AddWin persist a new Entry into mongodb fails collection.
func (client MongoAPIClient) AddFail() error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	collection := session.DB(client.dbClient.DbInfo.DB).C(client.dbClient.DbInfo.FailColl)
	err := collection.Insert(domain.NewEntry())
	if err != nil {
		log.Println(err)
	}
	return err
}
