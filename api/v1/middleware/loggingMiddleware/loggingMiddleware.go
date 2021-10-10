package loggingMiddleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// Do stuff here
		log.Println("************************************************")
		log.Println("Info: " + request.RequestURI)
		log.Println("Info: " + request.RemoteAddr)
		log.Println("************************************************")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(response, request)
	})
}
