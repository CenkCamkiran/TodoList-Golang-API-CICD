![](https://kaid.com.tr/wp-content/uploads/2021/02/gitlab-logo-gray-rgb.jpg)

# TodoList-Golang-API

## Features

- Developed via Gorilla/Mux Golang Library
- Use MongoDB Database
- Easy to deploy
- Included Unit Tests
- Can run on any platform (Mac, Linux ve Windows)
- Golang HTTP API


## Installation

```bash
$ git clone http://35.242.218.42/root/todolist-golang-api.git
$ cd todolist-golang-api/
$ go mod vendor
$ go run todo_api.go
# API Endpoint : http://127.0.0.1:9000
```

After download or clone project, run `go mod vendor` command. After the libraries installed, you should run `go run todo_api.go` command. TodoList API will listen on `9000` port.

  
## Structure
```
└───todolist-golang-api
    │   .gitlab-ci.yml // Gitlab file for our CI/CD pipeline
    │   Dockerfile // Dockerfile to build image
    │   go.mod
    │   go.sum
    │   README.md
    │   todo_api.go // main go file
    │
    ├───api
        └───v1
            ├───constants
            │       dbCredentials.go // Username and Password Credentials of MongoDB 
            │       routeList.go // Our API Route Names      
            │
            ├───controllers
            │   └───todoListController
            │           getTodoListController.go // GET /api/v1/getTodoList API Controller
            │           postTodoItemController.go // POST /api/v1/addNewTodoItem API Controller
            │           TodoListController_test.go // Unit test for Todolist Controllers
            │           
            │
            ├───dataAccess
            │   └───todoListDataAccess
            │           AddNewTodoItemToMongoDB.go // Data Access Layer to add new TodoList Item
            │           GetTodoListFromMongoDB.go // Data Access Layer to get TodoList
            │           todoListDataAccess_test.go // Unit test for Data Access Layer
            │
            ├───middleware
            │   └───loggingMiddleware
            │           loggingMiddleware.go // API Logging Middleware
            │
            ├───models
            │       TodoListModel.go // Models for our application
            │
            └───services
                ├───dbConnect
                │       mongoDBConnect.go // Service to connect MongoDB
                │       mongoDBConnect_test.go // Unit test for connection of MongoDB
                │
                └───responseFunctions
                │       responseFunctions.go // Service to response any request from outside
                │       responseFunctions_test.go // Unit test for response service
                │        
                └───todoListServices
                        getTodoList.go // Service Layer for our /api/v1/getTodoList API Route
                        addNewTodoItem.go // Service Layer for our /api/v1/addNewTodoItem API Route
                        todoListServices_test.go // Unit test for service layer
     
```


## API

#### /api/v1/getTodoList
* `GET` : Get all TodoList Items. You should use this route to get all TodoList Items from MongoDB. Response type is JSON.


#### /api/v1/addNewTodoItem
* `POST` : Add a new TodoList Item. You should use this route to get all TodoList Items from MongoDB. Response type is JSON.

```
{
    "Description": "Example TodoList Item"
}
```


## Test

```bash
run $ go test ./... -v -coverprofile .coverage.txt #and see the results in .coverage.txt file.

```


## Contributing

#### Bug Reports & Feature Requests

Please use the [issue tracker](http://35.242.218.42/root/todolist-golang-api/-/issues) to report any bugs or file feature requests.

#### Developing

PRs are welcome.