package web

import (
	"encoding/json"
	"net/http"
	io "playlist-generator/io"
	"playlist-generator/models"
	web "playlist-generator/web/actors"

	"github.com/mitchellh/mapstructure"
)

func UserHistoryHandler(writer http.ResponseWriter, request *http.Request) {
	var config = models.UserHistoryRequest{}
	genericConfig, err := io.UnmarshalGenericFunction(request.Body, config)

	if err != nil {
		handleError(err, writer)
		return
	}

	mapstructure.Decode(genericConfig, &config)

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
