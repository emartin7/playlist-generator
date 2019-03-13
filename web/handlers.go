package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"playlist-generator/errors"
	"playlist-generator/models"
	web "playlist-generator/web/actors"
)

func userHistoryHandler(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, bytesError := ioutil.ReadAll(request.Body)
	if bytesError != nil {
		handleError(&errors.UnmarshalError{Err: bytesError.Error()}, writer)
		return
	}

	var config models.UserHistoryRequest
	unmarshallingError := json.Unmarshal(bodyBytes, &config)
	if unmarshallingError != nil {
		handleError(&errors.UnmarshalError{Err: unmarshallingError.Error()}, writer)
		return
	}

	userHistoryRequestValidationError := models.ValidateUserHistoryRequest(config)
	if userHistoryRequestValidationError != nil {
		handleError(userHistoryRequestValidationError, writer)
		return
	}

	var requestError error
	var user models.PagingType

	if config.TypeOfSearch == "tracks" {
		user, requestError = web.GetUserHistoryTracks(config)
	} else {
		user, requestError = web.GetUserHistoryArtists(config)
	}

	if requestError != nil {
		handleError(requestError, writer)
		return
	}

	bytesOut, marshallingError := json.Marshal(user)
	if requestError != nil {
		handleError(marshallingError, writer)
		return
	}

	writer.Write(bytesOut)
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
			http.Error(writer, `{ "message: NotAuthorized" }`, http.StatusUnauthorized)
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
