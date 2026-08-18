package database

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString ="mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"


var collection *mongo.Collection

func init(){
	//client option
    clientoption := options.Client().ApplyURI(connectionString)
	//connect to mongodb
	client,err := mongo.Connect(context.TODO(),clientoption)
	if err != nil{
		panic(err)
	}
	fmt.Println("Mongodb connection sucess")

	collection=client.Database(dbName).Collection(colName)
	//collection istance 
	fmt.Println("collection istance is ready")
}