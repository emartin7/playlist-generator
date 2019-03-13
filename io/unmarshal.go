package io

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"playlist-generator/errors"
	"playlist-generator/models"
)

func UnmarshalToArtistsPaging(responseBody io.ReadCloser) (*models.ArtistsPaging, error) {
	defer responseBody.Close()

	bodyBytes, bytesError := ioutil.ReadAll(responseBody)
	if bytesError != nil {
		return nil, &errors.UnmarshalError{Err: bytesError.Error()}
	}

	var artistsPaging models.ArtistsPaging
	err := json.Unmarshal(bodyBytes, &artistsPaging)
	if err != nil {
		return nil, &errors.UnmarshalError{Err: err.Error()}
	}
	return &artistsPaging, nil
}

func UnmarshalToTracksPaging(responseBody io.ReadCloser) (*models.TracksPaging, error) {
	defer responseBody.Close()

	bodyBytes, bytesError := ioutil.ReadAll(responseBody)
	if bytesError != nil {
		return nil, &errors.UnmarshalError{Err: bytesError.Error()}
	}

	var tracksPaging models.TracksPaging
	err := json.Unmarshal(bodyBytes, &tracksPaging)
	if err != nil {
		return nil, &errors.UnmarshalError{Err: err.Error()}
	}
	return &tracksPaging, nil
}

func UnmarshalToRecommendationsResponse(responseBody io.ReadCloser) (*models.RecommendationResponse, error) {
	defer responseBody.Close()

	bodyBytes, bytesError := ioutil.ReadAll(responseBody)
	if bytesError != nil {
		return nil, &errors.UnmarshalError{Err: bytesError.Error()}
	}

	var recommendationResponse models.RecommendationResponse
	err := json.Unmarshal(bodyBytes, &recommendationResponse)
	if err != nil {
		return nil, &errors.UnmarshalError{Err: err.Error()}
	}
	return &recommendationResponse, nil
}
