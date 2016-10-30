package mongo

import (
	"os"
	"time"
	"strings"
	"log"

	"gopkg.in/mgo.v2"
	"github.com/joho/godotenv"
	"github.com/MarcosSegovia/MyWins/src/wins/domain"
)

var (
	client *MongoClient
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

type MongoClient struct {
	Session *mgo.Session
	DbInfo  *DbInfo
}

type DbInfo struct {
	DB       string
	WinColl  string
	FailColl string
}

func NewMongoClient() *MongoClient {

	if client != nil {

		return client
	}

	godotenv.Load()
	mongoDBHost := os.Getenv("DB_PORT_27017_TCP_ADDR")
	mongoDBPort := os.Getenv("DB_PORT_27017_TCP_PORT")
	mywinsDatabase := os.Getenv("DB_DBNAME")
	winsCollection := os.Getenv("DB_WINS_COLLECTION")
	failsCollection := os.Getenv("DB_FAILS_COLLECTION")
	timeout := time.Duration(40 * time.Second)
	host := []string{mongoDBHost}
	host = append(host, ":")
	host = append(host, mongoDBPort)
	hostName := strings.Join(host, "")
	mongoSession, err := mgo.DialWithTimeout(hostName, timeout)

	if err != nil {
		log.Printf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	client := &MongoClient{
		Session: mongoSession,
		DbInfo: &DbInfo{mywinsDatabase, winsCollection, failsCollection},
	}
	return client
}
