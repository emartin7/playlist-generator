package web

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	io "playlist-generator/io"
	"playlist-generator/models"
	web "playlist-generator/web/actors"

	"github.com/mitchellh/mapstructure"
)

func RecommendationHandler(writer http.ResponseWriter, request *http.Request) {
	defer recoverFromPanic(writer)
	var recommendationRequest = models.RecommendationRequest{}
	genericConfig, err := io.UnmarshalGenericFunction(request.Body, recommendationRequest)

	if err != nil {
		handleError(err, writer)
		return
	}

	mapstructure.Decode(genericConfig, &recommendationRequest)

	if recommendationRequest.UseUserHistory == true {
		recommendationRequest, err = useHistoryToGetSeeds(recommendationRequest)
		if err != nil {
			handleError(err, writer)
			return
		}
	}

	recommendationResponse, err := web.GetRecommendations(recommendationRequest)

	if err != nil {
		handleError(err, writer)
		return
	}

	bytesOut, err := json.Marshal(recommendationResponse)
	if err != nil {
		handleError(err, writer)
		return
	}

	writer.Write(bytesOut)
}

func getUserHistory(config models.UserHistoryRequest) (user models.PagingType, err error) {
	// err = models.ValidateUserHistoryRequest(config)
	if err != nil {
		log.Println("in the error statement")
		return
	}

	if config.TypeOfSearch == "tracks" {
		user, err = web.GetUserHistoryTracks(config)
		log.Println(user)
	} else {
		user, err = web.GetUserHistoryArtists(config)
	}

	return
}

func seedRecommendationObjectFromHistory(userHistory models.PagingType, recommendationRequest models.RecommendationRequest) (recommendation models.RecommendationRequest, err error) {
	var seed *models.Seeds

	if recommendationRequest.TypeOfSearch == "tracks" {
		seed, err = getSeedsFromTracks(userHistory.(*models.TracksPaging))
	} else {
		seed, err = getSeedsFromArtists(userHistory.(*models.ArtistsPaging))
	}
	recommendationRequest.SeedArtists = seed.SeedArtists
	recommendationRequest.SeedGenres = seed.SeedGenres
	recommendationRequest.SeedTracks = seed.SeedTracks

	return recommendationRequest, err
}

func getSeedsFromArtists(artists *models.ArtistsPaging) (seeds *models.Seeds, err error) {
	var artistSlice []string
	if len(artists.Items) < 1 {
		return nil, errors.New("Invalid Request: no history to search from")
	}
	for i := 0; i < len(artists.Items); i++ {
		artistSlice = append(artistSlice, artists.Items[i].ID)
	}

	if len(artistSlice) < 5 {
		seeds.SeedArtists = artistSlice
		return
	}

	for i := 0; i < 5; i++ {
		seeds.SeedArtists = append(seeds.SeedArtists, artistSlice[rand.Intn(len(artistSlice)-1)])
	}
	return
}

func getSeedsFromTracks(tracks *models.TracksPaging) (seeds *models.Seeds, err error) {
	log.Println(tracks)
	var tracksSlice []string
	if len(tracks.Items) < 1 {
		return nil, errors.New("Invalid Request: no history to search from")
	}
	for i := 0; i < len(tracks.Items); i++ {
		tracksSlice = append(tracksSlice, tracks.Items[i].ID)
	}

	if len(tracksSlice) < 5 {
		seeds.SeedTracks = tracksSlice
		return
	}

	for i := 0; i < 5; i++ {
		seeds.SeedTracks = append(seeds.SeedTracks, tracksSlice[rand.Intn(len(tracksSlice)-1)])
	}
	return
}

func useHistoryToGetSeeds(recommendationRequest models.RecommendationRequest) (recommendation models.RecommendationRequest, err error) {
	userHistoryRequest := models.UserHistoryRequest{
		TimeRange:    recommendationRequest.TimeRange,
		TypeOfSearch: recommendationRequest.TypeOfSearch,
		Offset:       recommendationRequest.Offset,
		Limit:        recommendationRequest.Limit,
		OauthToken:   recommendationRequest.OauthToken,
	}

	userHistory, err := getUserHistory(userHistoryRequest)
	log.Println(err)
	if err != nil {
		return
	}

	return seedRecommendationObjectFromHistory(userHistory, recommendationRequest)
}
