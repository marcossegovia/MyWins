package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"strings"
	"time"
)

const (
	TOKEN_DB_NAME = "TokenSessions"
	ACCESS_COL    = "accesses"
	REFRESH_TOKEN = "refreshtoken"
)

var (
	client             *MongoClient
	mongoConfiguration *MongoConfiguration
)

// MongoConfiguration holds all the needed configuration to establish connection with MongoDB.
type MongoConfiguration struct {
	MongoHost       string
	MongoPort       string
	MyWinsDatabase  string
	WinsCollection  string
	FailsCollection string
}

// MongoClient holds the Current Client Session and Database Information to interact with MongoDB.
type MongoClient struct {
	Session *mgo.Session
	DbInfo  *dbInfo
}

type dbInfo struct {
	DB       string
	WinColl  string
	FailColl string
}

// InitMongoEnv is called from the Environment configuration to set the connections parameters for MongoDB instance.
func InitMongoEnv(configuration *MongoConfiguration) {
	mongoConfiguration = configuration
}

// NewMongoClient instantiates the MongoDB Connection using the predefined configuration when setting the Environment (Production or Development).
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
		DbInfo:  &dbInfo{mongoConfiguration.MyWinsDatabase, mongoConfiguration.WinsCollection, mongoConfiguration.FailsCollection},
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
