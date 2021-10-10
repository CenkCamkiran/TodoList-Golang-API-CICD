package todoListDataAccess

import (
	"context"
	"net/http"
	"time"
	dbCredentials "todo_api/api/v1/constants"
	models "todo_api/api/v1/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTodoListFromMongoDB(mongoClient *mongo.Client) ([]models.TodoList, models.TodoListError) {

	var DBError models.TodoListError = models.TodoListError{
		Message:    "",
		StatusCode: 200,
	}

	MongoCollection := mongoClient.Database(dbCredentials.MONGODB_DATABASE).Collection(dbCredentials.MONGODB_TODOLIST_COLLECTION)

	var TodoList []models.TodoList

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//****************************************************************************

	cursor, err := MongoCollection.Find(ctx, bson.M{})
	if err != nil {

		DBError := models.TodoListError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}

		return nil, DBError

	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var TodoListItem models.TodoList
		cursor.Decode(&TodoListItem)
		TodoList = append(TodoList, TodoListItem)
	}

	if err := cursor.Err(); err != nil {

		DBError := models.TodoListError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}

		return nil, DBError

	}

	if len(TodoList) == 0 {

		DBError := models.TodoListError{
			Message:    "No data found.",
			StatusCode: http.StatusNoContent,
		}

		return nil, DBError
	}

	return TodoList, DBError

}
