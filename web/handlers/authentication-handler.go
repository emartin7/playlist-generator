package web

import (
	"log"
	"net/http"
	actors "playlist-generator/web/actors"
)

const clientId = "96f6ab889f904ec49f7e4e8bf04db1e8"
const authorizedPath = "/authorize"

func AuthenticationHandler(writer http.ResponseWriter, request *http.Request) {
	req, _ := actors.GetAuthorizedUserRequest(request)
	log.Println(req.URL.RawQuery)
	http.Redirect(writer, req, "https://accounts.spotify.com/authorize?"+req.URL.RawQuery, 302)
}

func AuthenticationRedirectHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Got Here")
}
