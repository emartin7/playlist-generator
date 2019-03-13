package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user-history", userHistoryHandler).Methods("POST")
	router.HandleFunc("/recommendations", recommendationHandler).Methods("POST")
	router.Use(loggingMiddleware)

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request URI: " + request.RequestURI)
		next.ServeHTTP(writer, request)
	})
}
