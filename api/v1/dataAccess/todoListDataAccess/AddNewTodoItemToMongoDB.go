package todoListDataAccess

import (
	"context"
	"log"
	"net/http"
	"time"
	dbCredentials "todo_api/api/v1/constants"
	models "todo_api/api/v1/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddNewTodoItemToMongoDB(mongoClient *mongo.Client, Description string) models.TodoListError {

	var MongoResult models.TodoListError = models.TodoListError{
		Message:    "TodoItem is successfully added!",
		StatusCode: 200,
	}

	MongoCollection := mongoClient.Database(dbCredentials.MONGODB_DATABASE).Collection(dbCredentials.MONGODB_TODOLIST_COLLECTION)

	//****************************************************************************

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	TodoItem := bson.D{
		{"Description", Description},
	}

	insertResult, err := MongoCollection.InsertOne(ctx, TodoItem)
	if err != nil {

		MongoResult := models.TodoListError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
		log.Println(err.Error())

		return MongoResult

	}

	log.Println("Info: "+" TodoItem is successfully added!: ", insertResult.InsertedID)

	return MongoResult

}
