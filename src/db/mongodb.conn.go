package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "TDTxLE"
	pwd      = "comemierda1"
	host     = "TDTxLE-mongodb"
	port     = 27017
	database = "Libreria"
)

type Mongodb struct {
	client *mongo.Client
}

func (db *Mongodb) GetCollection(coll string) *mongo.Collection {
	// uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		getenv("DB_USER", "TDTxLE"),
		getenv("DB_PWD", "comemierda1"),
		getenv("DB_HOST", ""),
		getenv("DB_POST", ""),
	)

	var err error

	if db.client, err = mongo.NewClient(options.Client().ApplyURI(uri)); err == nil {
		ctx, c := context.WithTimeout(context.Background(), 10*time.Second)
		defer c()
		if err = db.client.Connect(ctx); err != nil {
			panic(err.Error())
		}

		return db.client.Database(database).Collection(coll)
	} else {

		panic(err.Error())
	}

}

func (db *Mongodb) Close() {
	db.client.Disconnect(context.Background())
}

func getenv(key, defaultValue string) string {
	value, defined := os.LookupEnv(key)

	if !defined {
		return defaultValue
	}

	return value
}
