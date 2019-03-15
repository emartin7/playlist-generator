package models

type ClientRecommendationResponse struct {
	Tracks []string `json:"uris"`
}