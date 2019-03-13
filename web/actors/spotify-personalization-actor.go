package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/errors"
	"playlist-generator/models"
	"strconv"
)

const spotifyBaseAddress = "https://api.spotify.com"
const userHistoryPath = "/v1/me/top/"

func GetUserHistory(config models.UserHistoryRequest) (*models.ArtistsPaging, error) {
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

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		var user models.ArtistsPaging
		error := json.Unmarshal(bodyBytes, &user)
		if error != nil {
			return nil, &errors.UnmarshalError{Err: err.Error()}
		}
		return &user, nil
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: "Error in request to Spotify"}
}
