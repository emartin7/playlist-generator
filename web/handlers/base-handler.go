package web

import (
	"runtime/debug"
	"log"
	"net/http"
	"playlist-generator/errors"
	golangErrors "errors"
)

func recoverFromPanic(writer http.ResponseWriter) {
	if r := recover(); r != nil {
		log.Println("Recovered from: ", r)
		log.Println("Stacktrace from panic: \n" + string(debug.Stack()))
		handleError(golangErrors.New("System Failure"), writer)
	}
}

func handleError(err error, writer http.ResponseWriter) {
	log.Println(err.Error())

	switch err.(type) {
	case *errors.ValidationError:
		http.Error(writer, `{ "message: Invalid Request Body" }`, http.StatusBadRequest)
	case *errors.HttpError:
		switch (err.(*errors.HttpError)).StatusCode {
		case 400:
			http.Error(writer, `{ "message: BadRequest" }`, http.StatusBadRequest)
		case 401:
			http.Error(writer, `{ "message: Unauthorized" }`, http.StatusUnauthorized)
		default:
			http.Error(writer, `{ "message: ServiceUnavailable" }`, http.StatusServiceUnavailable)
		}
	case *errors.UnmarshalError:
		http.Error(writer, `{ "message: BadRequest" }`, http.StatusBadRequest)
	case *errors.MarshalError:
		http.Error(writer, `{ "message: BadRequest" }`, http.StatusBadRequest)
	default:
		http.Error(writer, `{ "message: ServiceUnavailable" }`, http.StatusServiceUnavailable)
	}
}
