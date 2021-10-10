package todoListController

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_api/api/v1/constants"
	models "todo_api/api/v1/models"
	mongoDBConnect "todo_api/api/v1/services/dbConnect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetTodoListController(t *testing.T) {

	t.Run("can fetch todo list items", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()

		request := getTodoListRequest()
		response := httptest.NewRecorder()

		// GetTodoListController(mongoClient, response, request)

		handler := http.HandlerFunc(handleRequest(mongoClient, GetTodoListController))
		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

	})

	t.Run("can fetch empty todo list", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()

		database := mongoClient.Database(constants.MONGODB_DATABASE)
		collection := database.Collection(constants.MONGODB_TODOLIST_COLLECTION)
		result, _ := collection.DeleteMany(context.TODO(), bson.M{})
		fmt.Printf("DeleteMany removed %v document(s) \n", result.DeletedCount)

		request := getTodoListRequest()
		response := httptest.NewRecorder()

		// GetTodoListController(mongoClient, response, request)

		handler := http.HandlerFunc(handleRequest(mongoClient, GetTodoListController))
		handler.ServeHTTP(response, request)

		var TodoListError models.TodoListError
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&TodoListError); err != nil {
			//
		}

		assertStatus(t, response.Code, http.StatusNoContent)
		assertResponseBody(t, TodoListError.Message, "No data found.")

	})

	t.Run("testing on internal server error while getting todolist", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		mongoClient.Disconnect(context.TODO())

		request := getTodoListRequest()
		response := httptest.NewRecorder()

		// GetTodoListController(mongoClient, response, request)

		handler := http.HandlerFunc(handleRequest(mongoClient, GetTodoListController))
		handler.ServeHTTP(response, request)

		var TodoListError models.TodoListError
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&TodoListError); err != nil {
			//
		}

		assertStatus(t, response.Code, http.StatusInternalServerError)
		assertResponseBody(t, fmt.Sprintf("%d", TodoListError.StatusCode), fmt.Sprintf("%d", http.StatusInternalServerError))

	})

}

func TestAddNewTodoItemController(t *testing.T) {

	t.Run("can add new todo item to mongodb", func(t *testing.T) {
		TodoItem := models.TodoListRequest{
			Description: "Testing...",
		}
		byteJSON, _ := json.MarshalIndent(TodoItem, "", "  ")

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()

		request := newTodoItemRequest(bytes.NewBuffer(byteJSON))
		response := httptest.NewRecorder()

		// AddNewTodoItemController(mongoClient, response, request)

		handler := http.HandlerFunc(handleRequest(mongoClient, AddNewTodoItemController))
		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		var NewTodoResponse models.TodoListError
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&NewTodoResponse); err != nil {
			//
		}

		assertResponseBody(t, NewTodoResponse.Message, "TodoItem is successfully added!")

	})

	t.Run("testing on internal server error while adding new todo item", func(t *testing.T) {

		TodoItem := models.TodoListRequest{
			Description: "Testing...",
		}
		byteJSON, _ := json.MarshalIndent(TodoItem, "", "  ")

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		mongoClient.Disconnect(context.TODO())

		request := newTodoItemRequest(bytes.NewBuffer(byteJSON))
		response := httptest.NewRecorder()

		// AddNewTodoItemController(mongoClient, response, request)

		handler := http.HandlerFunc(handleRequest(mongoClient, AddNewTodoItemController))
		handler.ServeHTTP(response, request)

		var NewTodoResponse models.TodoListError
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&NewTodoResponse); err != nil {
			//
		}

		assertStatus(t, response.Code, http.StatusInternalServerError)
		assertResponseBody(t, fmt.Sprintf("%d", NewTodoResponse.StatusCode), fmt.Sprintf("%d", http.StatusInternalServerError))

	})

	t.Run("testing wrong json body while adding new todo item", func(t *testing.T) {

		mongoClient, _ := mongoDBConnect.ConnectMongoDB()
		mongoClient.Disconnect(context.TODO())

		request := newTodoItemRequest(bytes.NewBuffer([]byte("{'username': 'download'email:'john@gmail.com'}")))
		response := httptest.NewRecorder()

		// AddNewTodoItemController(mongoClient, response, request)

		handler := http.HandlerFunc(handleRequest(mongoClient, AddNewTodoItemController))
		handler.ServeHTTP(response, request)

		var NewTodoResponse models.TodoListError
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&NewTodoResponse); err != nil {
			//
		}

		assertStatus(t, response.Code, http.StatusBadRequest)
		assertResponseBody(t, fmt.Sprintf("%d", NewTodoResponse.StatusCode), fmt.Sprintf("%d", http.StatusBadRequest))

	})

}

func handleRequest(client *mongo.Client, handler func(client *mongo.Client, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(client, w, r)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	} else {
		t.Logf("got correct status, got %d, want %d", got, want)
	}
}

func getTodoListRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/todoList", nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	} else {
		t.Logf("response body is correct, got %q want %q", got, want)
	}
}

func newTodoItemRequest(body *bytes.Buffer) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/todoList", body)
	return req
}
