package web

import (
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/errors"
	"playlist-generator/io"
	"playlist-generator/models"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const spotifyBaseAddress = "https://api.spotify.com"
const userHistoryPath = "/v1/me/top/"

func GetUserHistoryArtists(config models.UserHistoryRequest) (artists *models.ArtistsPaging, err error) {
	resp, err := getUserHistory(config)

	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		artistsContainer := models.ArtistsPaging{}
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, artistsContainer)
		mapstructure.Decode(genericResp, &artistsContainer)
		return &artistsContainer, err
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}

func GetUserHistoryTracks(config models.UserHistoryRequest) (tracks *models.TracksPaging, err error) {
	resp, err := getUserHistory(config)

	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		tracksContainer := models.TracksPaging{}
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, tracksContainer)
		mapstructure.Decode(genericResp, &tracksContainer)
		return &tracksContainer, err
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}

func getUserHistory(config models.UserHistoryRequest) (*http.Response, error) {
	resp, err := clients.Get(models.HttpRequest{
		QueryParams: map[string]string{
			"offset":     strconv.Itoa(config.Offset),
			"time_range": config.TimeRange,
			"limit":      strconv.Itoa(config.Limit),
		},
		Headers: map[string]string{
			"Authorization": "Bearer " + config.OauthToken,
		},
		Path: spotifyBaseAddress + userHistoryPath + config.TypeOfSearch,
	})

	if err != nil {
		return nil, &errors.HttpError{StatusCode: 503, Err: err.Error()}
	}

	return resp, nil
}
