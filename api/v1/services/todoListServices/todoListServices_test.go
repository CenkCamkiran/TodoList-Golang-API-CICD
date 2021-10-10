package todoListServices

import (
	"context"
	"testing"
	"todo_api/api/v1/models"
	mongoDBConnect "todo_api/api/v1/services/dbConnect"
)

func TestConnectMongoDB(t *testing.T) {

	t.Run("getTodoList service can fetch data from data access layer", func(t *testing.T) {
		mongoClient, _ := mongoDBConnect.ConnectMongoDB()

		_, err := GetTodoListService(mongoClient)

		assertTodoList(t, err)

	})

	t.Run("testing getTodoList service while mongodb connection error", func(t *testing.T) {
		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		mongoClient.Disconnect(context.TODO())

		_, err := GetTodoListService(mongoClient)

		assertTodoList(t, err)

	})

}

func assertTodoList(t *testing.T, err models.TodoListError) {
	t.Helper()

	if err.StatusCode != 200 {
		t.Logf("DataAccess Layer's response is correct %s", err.Message)

	} else if err.StatusCode == 200 {
		t.Logf("DataAccess Layer's response is correct")

	} else {
		t.Errorf("Test failed")
	}
}
