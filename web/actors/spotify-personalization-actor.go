package web

import (
	"github.com/mitchellh/mapstructure"
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/errors"
	"playlist-generator/io"
	"playlist-generator/models"
	"strconv"
)

const spotifyBaseAddress = "https://api.spotify.com"
const userHistoryPath = "/v1/me/top/"

func GetUserHistoryArtists(config models.UserHistoryRequest) (artists *models.ArtistsPaging, err error) {
	resp, err := getUserHistory(config)

	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, models.ArtistsPaging{})
		mapstructure.Decode(genericResp, &artists)
		return artists, err
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}

func GetUserHistoryTracks(config models.UserHistoryRequest) (tracks *models.TracksPaging, err error) {
	resp, err := getUserHistory(config)

	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, models.TracksPaging{})
		mapstructure.Decode(genericResp, &tracks)
		return tracks, err
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
