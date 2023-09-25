package config

import (
	"User/constants"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoCollection := options.Client().ApplyURI(constants.ConnectionString)
	mongoClient,err:=mongo.Connect(ctx,mongoCollection)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err 
	}
	// if err:= mongoClient.Ping(ctx,readpref.Primary());err!=nil{
	// 	return nil,err
	// }
	return mongoClient,err
}
func GetCollection(client *mongo.Client,dbname string,collectionName string)*mongo.Collection{
	collection := client.Database(dbname).Collection(collectionName)
	return collection
}