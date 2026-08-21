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

// error func
func Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

func InsertOnemovie(movie model.Netflix) {
	insert, err := collection.InsertOne(context.Background(), movie)
	Error(err)
	fmt.Println("Inserted One movie in db with id : ", insert.InsertedID)
}

func UpdateOnemovie(movieId string) {

	id, err := bson.ObjectIDFromHex(movieId)
	Error(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}

//delete one movie

func DeleteOnemovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)
	Error(err)
	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	Error(err)
	fmt.Println("Successfully deleted the movie with the count : ", deletecount)
}

//delete the complete database data

func DeleteAllmovie() {
	filter := bson.D{{}}
	//u could declare a var of filter or directly pass bson.D
	deletecount, err := collection.DeleteMany(context.Background(), filter, nil)
	Error(err)
	fmt.Println("Number of movies deleted is : ", deletecount)

}

//get all movies from db

func Getallmovies() []model.Netflix {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	Error(err)
	defer cur.Close(context.Background())
	var movies []model.Netflix
	for cur.Next(context.Background()) {
		var movie model.Netflix
		err := cur.Decode(&movie)
		Error(err)
		movies = append(movies, movie)
	}

	return movies
}
