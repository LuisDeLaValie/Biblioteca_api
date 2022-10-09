package db

import (
	
	"fmt"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	usr = "TDTxLE"
	pwd = "comemierda1"
	host = "TDTxLE-mongodb"
	port = 27017
	database = "Libreria"
)

func GetCollection(coll string) *mongo.Collection {
	// uri := os.Getenv("MONGODB_URI")
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",usr,pwd,host,port)
	// uri := fmt.Sprintf("mongodb://%s:%d",host,port)
	client,err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx,_ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(coll)

}