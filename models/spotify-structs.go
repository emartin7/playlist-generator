package models

type PagingType interface {
	GetType() string
}
type AlbumSimplified struct {
	AlbumGroup           *string             `json:"album_group"`
	AlbumType            *string             `json:"album_type"`
	Artists              []ArtistsSimplified `json:"artists"`
	AvailableMarkets     []string            `json:"available_markets"`
	ExternalUrls         ExternalUrls        `json:"external_urls"`
	Href                 *string             `json:"href"`
	ID                   *string             `json:"id"`
	Images               []Image             `json:"images"`
	Name                 *string             `json:"name"`
	ReleaseDate          *string             `json:"release_date"`
	ReleaseDatePrecision *string             `json:"release_date_precision"`
	Type                 *string             `json:"type"`
	URI                  *string             `json:"uri"`
}

type ArtistsPaging struct {
	Items []ArtistsSimplified `json:"items"`
	*Paging
}

type ArtistsFull struct {
	*ArtistsSimplified
	Followers  Followers `json:"followers"`
	Genres     *[]string `json:"genres"`
	Images     []Image   `json:"images"`
	Popularity int       `json:"popularity"`
}

type ArtistsSimplified struct {
	Href         *string      `json:"href"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	ID           string       `json:"id"`
	Name         *string      `json:"name"`
	Type         *string      `json:"type"`
	URI          *string      `json:"uri"`
}

type ExternalUrls struct {
	Spotify *string `json:"spotify"`
}

type Followers struct {
	Href  *string `json:"href"`
	Total int     `json:"total"`
}

type Image struct {
	Height int     `json:"height"`
	URL    *string `json:"url"`
	Width  int     `json:"width"`
}

type Paging struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Total    int     `json:"total"`
	Limit    int     `json:"limit"`
	Href     *string `json:"href"`
}

type RecommendationResponse struct {
	Seeds  []SeedRecommendations `json:"seeds"`
	Tracks []TracksFull          `json:"tracks"`
}

type SeedRecommendations struct {
	AfterFilteringSize int     `json:"afterFilteringSize"`
	AfterRelinkingSize int     `json:"afterRelinkingSize"`
	Href               *string `json:"href"`
	ID                 *string `json:"id"`
	InitialPoolSize    int     `json:"initialPoolSize"`
	Type               *string `json:"type"`
}

type TracksFull struct {
	Album            AlbumSimplified     `json:"album"`
	Artists          []ArtistsSimplified `json:"artists"`
	AvailableMarkets []string            `json:"available_markets"`
	DurationMs       int                 `json:"duration_ms"`
	Explicit         bool                `json:"explicit"`
	ExternalUrls     ExternalUrls        `json:"external_urls"`
	Href             *string             `json:"href"`
	ID               string              `json:"id"`
	IsPlayable       bool                `json:"is_playable"`
	Name             *string             `json:"name"`
	Popularity       int                 `json:"popularity"`
	Previewurl       *string             `json:"preview_url"`
	Type             *string             `json:"type"`
	URI              string              `json:"uri"`
}

type TracksPaging struct {
	Items []TracksFull `json:"items"`
	*Paging
}

/*
* Here is where implementations of the interfaces go
 */

func (artistPaging ArtistsPaging) GetType() string {
	return "artists"
}

func (trackPaging TracksPaging) GetType() string {
	return "tracks"
}
