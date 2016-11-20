package config

import (
	"os"
	"fmt"

	"github.com/MarcosSegovia/MyWins/src/wins/infrastructure/mongo"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func SetDevEnvironment() {
	viper.SetConfigName("mongo_dev")
	viper.AddConfigPath("$GOPATH/src/github.com/MarcosSegovia/MyWins/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	mongoDBHost := viper.GetString("mongoDBHost")
	mongoDBPort := viper.GetString("mongoDBPort")
	mywWnsDatabase := viper.GetString("mywinsDatabase")
	winsCollection := viper.GetString("winsCollection")
	failsCollection := viper.GetString("failsCollection")

	mongoConfiguration := &mongo.MongoConfiguration{
		MongoHost: mongoDBHost,
		MongoPort:mongoDBPort,
		MyWinsDatabase: mywWnsDatabase,
		WinsCollection: winsCollection,
		FailsCollection: failsCollection,
	}

	mongo.InitMongoEnv(mongoConfiguration)
}

func SetProdEnvironment() {

	godotenv.Load()
	mongoDBHost := os.Getenv("DB_PORT_27017_TCP_ADDR")
	mongoDBPort := os.Getenv("DB_PORT_27017_TCP_PORT")
	mywWnsDatabase := os.Getenv("DB_DBNAME")
	winsCollection := os.Getenv("DB_WINS_COLLECTION")
	failsCollection := os.Getenv("DB_FAILS_COLLECTION")

	mongoConfiguration := &mongo.MongoConfiguration{
		MongoHost: mongoDBHost,
		MongoPort:mongoDBPort,
		MyWinsDatabase: mywWnsDatabase,
		WinsCollection: winsCollection,
		FailsCollection: failsCollection,
	}

	mongo.InitMongoEnv(mongoConfiguration)
}
