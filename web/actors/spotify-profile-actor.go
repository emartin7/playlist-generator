package web

import (
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/errors"
	"playlist-generator/io"
	"playlist-generator/models"

	"github.com/mitchellh/mapstructure"
)

const profilePath = "/v1/recommendations"

func GetProfile(oauthToken string) (spotifyUser *models.SpotifyUser, err error) {
	resp, err := clients.Get(models.HttpRequest{
		Headers: map[string]string{
			"Authorization": "Bearer " + oauthToken,
		},
		Path: spotifyBaseAddress + profilePath,
	})

	if err != nil {
		return nil, &errors.HttpError{StatusCode: 503, Err: err.Error()}
	}

	if resp.StatusCode == http.StatusOK {
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, models.SpotifyUser{})
		mapstructure.Decode(genericResp, &spotifyUser)
		return spotifyUser, err
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}
