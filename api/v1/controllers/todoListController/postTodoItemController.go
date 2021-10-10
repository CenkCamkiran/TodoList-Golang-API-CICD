package todoListController

import (
	"encoding/json"
	"log"
	"net/http"
	models "todo_api/api/v1/models"
	"todo_api/api/v1/services/responseFunctions"
	addNewTodoItem "todo_api/api/v1/services/todoListServices"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddNewTodoItemController(mongoClient *mongo.Client, response http.ResponseWriter, request *http.Request) {

	var TodoListRequestModel models.TodoListRequest
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&TodoListRequestModel); err != nil {

		BadRequest := models.TodoListJSONError{
			Message:    "",
			StatusCode: http.StatusBadRequest,
		}

		byteJSON, _ := json.MarshalIndent(BadRequest, "", "  ")
		responseFunctions.RespondWithJSON(response, http.StatusBadRequest, byteJSON)

	}
	defer request.Body.Close()

	MongoResult := addNewTodoItem.AddNewTodoItemService(TodoListRequestModel.Description, mongoClient)

	log.Printf("%d", MongoResult.StatusCode)
	log.Printf("\n")
	switch MongoResult.StatusCode {
	case 200:

		byteJSON, err := json.MarshalIndent(MongoResult, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at addNewTodoItem API Endpoint " + err.Error())

			TodoListJSONError := models.TodoListJSONError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}

			byteJSON, _ := json.MarshalIndent(TodoListJSONError, "", "  ")
			responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

		}

		responseFunctions.RespondWithJSON(response, http.StatusOK, byteJSON)

	case 500:

		byteJSON, err := json.MarshalIndent(MongoResult, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at addNewTodoItem API Endpoint " + err.Error())

			TodoListJSONError := models.TodoListJSONError{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}

			byteJSON, _ := json.MarshalIndent(TodoListJSONError, "", "  ")
			responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

		}

		responseFunctions.RespondWithJSON(response, http.StatusInternalServerError, byteJSON)

	default:

		byteJSON, err := json.MarshalIndent(MongoResult, "", "  ")

		if err != nil {
			log.Println("Error: " + "Error at addNewTodoItem API Endpoint " + err.Error())

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
