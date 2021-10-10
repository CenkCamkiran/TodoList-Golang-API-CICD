package todoListServices

import (
	dataAccess "todo_api/api/v1/dataAccess/todoListDataAccess"
	models "todo_api/api/v1/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetTodoListService(mongoClient *mongo.Client) ([]models.TodoList, models.TodoListError) {

	TodoList, TodoListError := dataAccess.GetTodoListFromMongoDB(mongoClient)

	return TodoList, TodoListError

}
