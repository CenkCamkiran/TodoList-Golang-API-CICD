package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	routeList "todo_api/api/v1/constants"
	todoListController "todo_api/api/v1/controllers/todoListController"
	loggingMiddleware "todo_api/api/v1/middleware/loggingMiddleware"
	mongoDBConnect "todo_api/api/v1/services/dbConnect"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	MongoClient *mongo.Client
}

func ConfigAndRunApp() {
	app := new(App)
	apiStatus := app.Initialize()

	if apiStatus {
		app.Run("9000")

	} else {
		log.Println("Error: " + "API is not started! Check for errors!")

	}

}

func (app *App) Initialize() bool {

	mongoClient, err := mongoDBConnect.ConnectMongoDB()
	app.MongoClient = mongoClient

	if err != nil {

		log.Println("Error: " + err.Error())

		return false

	} else {
		app.Router = mux.NewRouter()
		app.UseMiddleware(loggingMiddleware.LoggingMiddleware)
		app.setRouters()

		return true
	}

}

func (app *App) setRouters() {
	app.Get(routeList.TODO_LIST_ROUTE, app.handleRequest(todoListController.GetTodoListController))
	app.Post(routeList.TODO_LIST_ROUTE, app.handleRequest(todoListController.AddNewTodoItemController))
}

func (app *App) UseMiddleware(middleware mux.MiddlewareFunc) {
	app.Router.Use(middleware)
}

func (app *App) Get(path string, endpoint http.HandlerFunc) {
	app.Router.HandleFunc(path, endpoint).Methods("GET")
}

func (app *App) Post(path string, endpoint http.HandlerFunc) {
	app.Router.HandleFunc(path, endpoint).Methods("POST")
}

func (app *App) Run(port string) {
	// use signals for shutdown server gracefully.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
		methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
		origins := handlers.AllowedOrigins([]string{"*"})

		log.Fatal(http.ListenAndServe(":9000", handlers.CORS(headers, methods, origins)(app.Router)))
	}()
	log.Printf("Server is listning on %s\n", port)
	sig := <-sigs
	log.Println("Signal: ", sig)

}

// RequestHandlerFunction is a custome type that help us to pass client arg to all endpoints
type RequestHandlerFunction func(client *mongo.Client, w http.ResponseWriter, r *http.Request)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.MongoClient, w, r)
	}
}

func main() {

	ConfigAndRunApp()

}
