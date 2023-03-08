package atdb

import (
	"context"
	"fmt"

	"github.com/aiteung/atmodel"
	"go.mongodb.org/mongo-driver/bson"
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

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("AIteung Mongo, InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetOneDoc(db *mongo.Database, collection string, filter bson.M) (notif atmodel.Notif) {
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&notif)
	if err != nil {
		fmt.Printf("GetNotifFromPhoneNumber: %v\n", err)
	}
	return
}

func DeleteOneDoc(db *mongo.Database, collection string, filter bson.M) (result *mongo.DeleteResult) {
	result, err := db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetNotifFromPhoneNumber: %v\n", err)
	}
	return
}
