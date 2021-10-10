package mongoDBConnect

import (
	"context"
	"fmt"
	"log"
	dbCredentials "todo_api/api/v1/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodbHost string = dbCredentials.MONGODB_HOST
var mongodbPort string = dbCredentials.MONGODB_PORT
var mongodbUsername string = dbCredentials.MONGODB_USERNAME
var mongodbPassword string = dbCredentials.MONGODB_PASSWORD
var mongodbDatabase string = dbCredentials.MONGODB_DATABASE
var mongodbAuthSource string = dbCredentials.MONGODB_AUTH_SOURCE

func ConnectMongoDB() (*mongo.Client, error) {
	mongoConnectionUrl := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=SCRAM-SHA-256&authSource=%s", mongodbUsername, mongodbPassword, mongodbHost, mongodbPort, mongodbDatabase, mongodbAuthSource)
	clientOptions := options.Client().ApplyURI(mongoConnectionUrl)

	client, MongoConnectionError := mongo.Connect(context.TODO(), clientOptions)

	if MongoConnectionError != nil {
		log.Println("Error: " + "MongoDB Ping Error!")

		return nil, MongoConnectionError
	}

	// Check the connection
	MongoConnectionError = client.Ping(context.TODO(), nil)

	if MongoConnectionError != nil {
		log.Println("Error: " + "MongoDB'ye ping atılamadı! Ping Error!")
		log.Println(MongoConnectionError.Error())

		return nil, MongoConnectionError
	}

	log.Println("Info: " + "MongoDB'ye bağlanıldı")

	return client, nil
}
