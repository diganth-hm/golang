package database

import (
	"context"
	"fmt"
	"log"

	model "github.com/diganth-hm/golang/buildapi/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientoption := options.Client().ApplyURI(connectionString)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongodb connection sucess")

	collection = client.Database(dbName).Collection(colName)
	//collection istance
	fmt.Println("collection istance is ready")
}

//mongodb helper file

func insertOnemovie(movie model.Netflix) {
	insert, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One movie in db with id : ", insert.InsertedID)
}

func updateOnemovie(movieId string) {

	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.updateOnemovie(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}
