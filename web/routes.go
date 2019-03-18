package web

import (
	"log"
	"net/http"
	web "playlist-generator/web/handlers"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	user := router.PathPrefix("/api/playlist-generator/user").Subrouter()
	user.HandleFunc("/user-history", web.UserHistoryHandler).Methods("POST")
	user.HandleFunc("/recommendations", web.RecommendationHandler).Methods("POST")
	user.Use(authenticationMiddleware)

	auth := router.PathPrefix("/api/playlist-generator/").Subrouter()
	auth.HandleFunc("/auth", web.UserHistoryHandler).Methods("POST")

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request URI: " + request.RequestURI)
		next.ServeHTTP(writer, request)
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("Authorization") == "" {
			http.Error(writer, http.StatusText(401), 401)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
