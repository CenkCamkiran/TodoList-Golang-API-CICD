package todoListServices

import (
	dataAccess "todo_api/api/v1/dataAccess/todoListDataAccess"
	models "todo_api/api/v1/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddNewTodoItemService(Description string, mongoClient *mongo.Client) models.TodoListError {

	MongoResult := dataAccess.AddNewTodoItemToMongoDB(mongoClient, Description)

	return MongoResult

}
