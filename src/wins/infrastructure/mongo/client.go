package mongo

import (
	"github.com/joho/godotenv"
	"os"
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
)

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

	accesses := mongoSession.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	index := defineIndex()

	err = accesses.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	client := &MongoClient{
		Session: mongoSession,
		DbInfo: &DbInfo{mywinsDatabase, winsCollection, failsCollection},
	}
	return client
}

func defineIndex() mgo.Index {

	return mgo.Index{
		Key:        []string{REFRESH_TOKEN},
		Unique:     false, // refreshtoken is sometimes empty
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
}
