package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MarcosSegovia/MyWins/src/wins/domain"
	"github.com/joho/godotenv"

	mgo "gopkg.in/mgo.v2"
)

var (
	mongoDBHosts string
	authDatabase string
	authUserName string
	authPassword string
	dbInfo       *domain.DBInfo
)

func init() {
	godotenv.Load()
	mongoDBHosts = os.Getenv("MONGO_HOST")
	authDatabase = os.Getenv("MONGO_AUTH_DATABASE")
	authUserName = os.Getenv("MONGO_USER")
	authPassword = os.Getenv("MONGO_PASS")
	mywinsDatabase := os.Getenv("MONGO_MYWINS_DATABASE")
	winsCollection := os.Getenv("MONGO_WINS_COLLECTION")
	failsCollection := os.Getenv("MONGO_FAILS_COLLECTION")
	dbInfo = domain.NewDBInfo(mywinsDatabase, winsCollection, failsCollection)
}

func main() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  60 * time.Second,
		Database: authDatabase,
		Username: authUserName,
		Password: authPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	router := NewRouter(mongoSession, dbInfo)
	e := http.ListenAndServe(":8080", router)

	if e != nil {

		log.Fatal(e)
	}
}
