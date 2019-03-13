package models

import "playlist-generator/errors"

type UserHistoryRequest struct {
	TypeOfSearch string `json:"typeOfSearch"`
	TimeRange    string `json:"timeRange"`
	Offset       int    `json:"offset"`
	Limit        int    `json:"limit"`
	OauthToken   string `json:"oauthToken"`
}

func ValidateUserHistoryRequest(request UserHistoryRequest) *errors.ValidationError {
	errorSlice := []string{}
	if request.TypeOfSearch == "" {
		errorSlice = append(errorSlice, "TypeOfSearch was empty")
	}
	if request.TypeOfSearch != "artists" && request.TypeOfSearch != "tracks" {
		errorSlice = append(errorSlice, "TypeOfSearch must be either artists or tracks")
	}
	if request.TimeRange != "short_term" && request.TimeRange != "medium_term" && request.TimeRange != "long_term" {
		errorSlice = append(errorSlice, "TimeRange parameter was invalid")
	}
	if request.OauthToken == "" {
		errorSlice = append(errorSlice, "OauthToken was empty")
	}
	if len(errorSlice) != 0 {
		validationError := errors.ValidationError{
			Err:                "UserHistoryRequest_Validation_Error",
			ProblemDiscriptors: errorSlice,
		}
		return &validationError
	}
	return nil
}
