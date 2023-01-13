package atdb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn DBInfo) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
	if err != nil {
		fmt.Printf("AIteung Mongo, MongoConnect: %v\n", err)
	}
	return client.Database(mconn.DBName)
}

func InsertOneDoc(mconn DBInfo, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(mconn).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("AIteung Mongo, InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
