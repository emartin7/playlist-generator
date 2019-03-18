package web

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/errors"
	"playlist-generator/io"
	"playlist-generator/models"
	"strings"

	"github.com/mitchellh/mapstructure"
)

var playlistCreatePath = "/v1/users/~/playlists"
var playlistAppendTrackPath = "/v1/playlists/~/tracks"

func CreatePlaylist(playlist models.SpotifyPlaylist, user models.SpotifyUser, oauthToken string) (playlist *models.SpotifyPlaylist, err error) {
	playlistBytes, _ := json.Marshal(playlist)
	resp, err := clients.Post(models.HttpRequest{
		Headers: map[string]string{
			"Authorization": "Bearer " + oauthToken,
		},
		Body: ioutil.NopCloser(bytes.NewReader(playlistBytes)),
		Path: spotifyBaseAddress + strings.Replace(playlistCreatePath, "~", user.ID, 1),
	})

	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		playlistContainer := models.SpotifyPlaylist{}
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, playlistContainer)
		mapstructure.Decode(genericResp, &playlistContainer)
		return &playlistContainer, err
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}

func AppendTracksToPlaylist(playlist models.SpotifyPlaylist, tracks models.RecommendationResponse, oauthToken string) (ok bool, err error) {
	trackBytes, _ := json.Marshal(tracks)
	resp, err := clients.Post(models.HttpRequest{
		Headers: map[string]string{
			"Authorization": "Bearer " + oauthToken,
		},
		Path: spotifyBaseAddress + strings.Replace(playlistAppendTrackPath, "~", playlist.ID, 1),
		Body: ioutil.NopCloser(bytes.NewReader(trackBytes)),
	})

	if err != nil {
		return
	}

	if resp.StatusCode == http.StatusOK {
		return true, err
	}
	return false, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}
