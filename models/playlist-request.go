package models

type PlaylistRequest struct {
	*ClientRecommendationResponse
	*SpotifyPlaylist
}