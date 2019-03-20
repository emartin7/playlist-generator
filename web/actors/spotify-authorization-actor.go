package web

import (
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/models"
)

const authorizePath = "/authorize"
const clientId = "96f6ab889f904ec49f7e4e8bf04db1e8"

func GetAuthorizedUserRequest(request *http.Request) (*http.Request, error) {
	req, err := clients.GetFullRequest("GET", models.HttpRequest{
		QueryParams: map[string]string{
			"client_id":     clientId,
			"response_type": "token",
			"redirect_uri":  "http://10.1.148.137/api/playlist-generator/test",
			"scopes":        "user-read-private user-read-email",
		},
		Path: SpotifyBaseAddress + authorizePath,
	})
	return req, err
}
