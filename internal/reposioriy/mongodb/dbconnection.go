package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CreateMongoClient() (*mongo.Client, error) {

	uri := os.Getenv("MONGO_URI")
	// uri := "mongodb+srv://bhavyashreehm_db_user:XE9qMLOtFbNC3WUQ@cluster0.zz6aclj.mongodb.net/?appName=Cluster0"
	// fmt.Println(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Println(err)
		panic(err)
	}else{
		fmt.Println("MongoDB Connected succesfully")
	}
	return client, nil

}
