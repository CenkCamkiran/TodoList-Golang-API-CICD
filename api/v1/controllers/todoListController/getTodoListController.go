package todoListController

import (
	"encoding/json"
	"log"
	"net/http"
	models "todo_api/api/v1/models"
	"todo_api/api/v1/services/responseFunctions"
	getTodoList "todo_api/api/v1/services/todoListServices"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetTodoListController(mongoClient *mongo.Client, response http.ResponseWriter, request *http.Request) {

	TodoList, DBError := getTodoList.GetTodoListService(mongoClient)
	log.Println(DBError.StatusCode)

	switch DBError.StatusCode {

	case 200:

		byteJSON, err := json.MarshalIndent(TodoList, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at getTodoList API Endpoint " + err.Error())

			todoListJSONError := models.TodoListJSONError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}

			byteJSON, _ := json.MarshalIndent(todoListJSONError, "", "  ")
			responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

		}

		responseFunctions.RespondWithJSON(response, http.StatusOK, byteJSON)

	case 204:

		byteJSON, err := json.MarshalIndent(DBError, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at getTodoList API Endpoint " + err.Error())

			todoListJSONError := models.TodoListJSONError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}

			byteJSON, _ := json.MarshalIndent(todoListJSONError, "", "  ")
			responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

		}

		responseFunctions.RespondWithJSON(response, http.StatusNoContent, byteJSON)

	case 500:

		byteJSON, err := json.MarshalIndent(DBError, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at getTodoList API Endpoint " + err.Error())

			todoListJSONError := models.TodoListJSONError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}

			byteJSON, _ := json.MarshalIndent(todoListJSONError, "", "  ")
			responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

		}

		responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

	default:

		byteJSON, err := json.MarshalIndent(DBError, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at getTodoList API Endpoint " + err.Error())

			TodoListJSONError := models.TodoListJSONError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}

			byteJSON, _ := json.MarshalIndent(TodoListJSONError, "", "  ")
			responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

		}

		responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)
	}

}
