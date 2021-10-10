package todoListDataAccess

import (
	"context"
	"net/http"
	"testing"
	mongoDBConnect "todo_api/api/v1/services/dbConnect"
)

func TestAddNewTodoItemToMongoDB(t *testing.T) {

	t.Run("testing data access layer for adding new todo item", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		TodoListError := AddNewTodoItemToMongoDB(mongoClient, "Testing Golang")

		assertNewTodoItem(t, TodoListError.StatusCode, http.StatusOK)

	})

	t.Run("testing data access layer for adding new todo item while mongodb connection error", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		mongoClient.Disconnect(context.TODO())

		TodoListError := AddNewTodoItemToMongoDB(mongoClient, "Testing Golang")

		assertNewTodoItem(t, TodoListError.StatusCode, http.StatusInternalServerError)

	})

}

func TesGetTodoListFromMongoDB(t *testing.T) {

	t.Run("testing data access layer for getting todo list", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		_, Err := GetTodoListFromMongoDB(mongoClient)

		assertTodoList(t, Err.StatusCode, http.StatusOK)

	})

	t.Run("testing data access layer for getting todo list while mongodb connection error", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		mongoClient.Disconnect(context.TODO())

		_, Err := GetTodoListFromMongoDB(mongoClient)

		assertTodoList(t, Err.StatusCode, http.StatusOK)

	})

}

func assertNewTodoItem(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("DataAccess Layer's response is wrong, got %q want %q", got, want)
	} else {
		t.Logf("DataAccess Layer's response is correct, got %q want %q", got, want)
	}
}

func assertTodoList(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("DataAccess Layer's response is wrong, got %q want %q", got, want)
	} else {
		t.Logf("DataAccess Layer's response is correct, got %q want %q", got, want)
	}
}
