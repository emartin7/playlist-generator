package models

type ArtistsPaging struct {
	Items  []ArtistFull `json:"items"`
	Paging Paging
}

type ArtistFull struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       *[]string    `json:"genres"`
	Href         *string      `json:"href"`
	ID           *string      `json:"id"`
	Images       []Image      `json:"images"`
	Name         *string      `json:"name"`
	Popularity   int          `json:"popularity"`
	Type         *string      `json:"type"`
	URI          *string      `json:"uri"`
}

type ArtistSimplified struct {
	Href         *string      `json:"href"`
	ExternalUrls ExternalUrls `json:"external_urls"`
	ID           *string      `json:"id"`
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
