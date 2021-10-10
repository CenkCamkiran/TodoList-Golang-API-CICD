package models

type TodoList struct {
	Description string `json:"Description" bson:"Description"`
}

type TodoListRequest struct {
	Description string `json:"Description"`
}

type TodoListError struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}

type TodoListJSONError struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}

type TodoListNoRecord struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}
