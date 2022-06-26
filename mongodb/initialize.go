package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var ctx context.Context

func FindAll(collection string) []bson.M {
	initConnection()
	defer db.Client().Disconnect(ctx)

	cursor, err := db.Collection(collection).Find(ctx, bson.M{})
	if err != nil {
		log.Println("PANIC 1")
		defer cursor.Close(ctx)
		panic(err)
	} else {
		var results []bson.M
		err = cursor.All(ctx, &results)
		if err != nil {
			log.Println("PANIC 2")
			panic(err)
		}
		return results
	}
}

func initConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI_STRING")))
	if err != nil {
		log.Println(err)
	}

	ctx = context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
	}

	db = client.Database(os.Getenv("MONGODB_DB_NAME"))
	if err != nil {
		log.Println(err)
	}
}
