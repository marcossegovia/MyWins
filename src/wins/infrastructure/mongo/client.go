package mongo

import (
	"time"
	"strings"
	"gopkg.in/mgo.v2"
	"log"
)

const (
	TOKEN_DB_NAME = "TokenSessions"
	ACCESS_COL = "accesses"
	REFRESH_TOKEN = "refreshtoken"
)

var (
	client *MongoClient
	mongoConfiguration *MongoConfiguration
)

type MongoConfiguration struct {
	MongoHost       string
	MongoPort       string
	MyWinsDatabase  string
	WinsCollection  string
	FailsCollection string
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

func InitMongoEnv(configuration *MongoConfiguration) {
	mongoConfiguration = configuration
}

func NewMongoClient() *MongoClient {

	if client != nil {

		return client
	}

	timeout := time.Duration(40 * time.Second)
	host := []string{mongoConfiguration.MongoHost}
	host = append(host, ":")
	host = append(host, mongoConfiguration.MongoPort)
	hostName := strings.Join(host, "")
	mongoSession, err := mgo.DialWithTimeout(hostName, timeout)

	if err != nil {
		log.Printf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	accesses := mongoSession.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	index := defineIndex()

	err = accesses.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	client := &MongoClient{
		Session: mongoSession,
		DbInfo: &DbInfo{mongoConfiguration.MyWinsDatabase, mongoConfiguration.WinsCollection, mongoConfiguration.FailsCollection},
	}
	return client
}

func defineIndex() mgo.Index {

	return mgo.Index{
		Key:        []string{REFRESH_TOKEN},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
}
