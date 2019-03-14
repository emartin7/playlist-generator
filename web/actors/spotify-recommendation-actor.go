package web

import (
	"github.com/mitchellh/mapstructure"
	"net/http"
	"playlist-generator/clients"
	"playlist-generator/errors"
	"playlist-generator/io"
	"playlist-generator/models"
	"strconv"
	"strings"
)

const recommendationPath = "/v1/recommendations"

func GetRecommendations(config models.RecommendationRequest) (recommendationResponse *models.RecommendationResponse, err error) {
	resp, err := clients.Get(models.HttpRequest{
		QueryParams: map[string]string{
			"limit":                   strconv.Itoa(config.Limit),
			"seed_artists":            strings.Join(config.SeedArtists, ","),
			"seed_genres":             strings.Join(config.SeedGenres, ","),
			"seed_tracks":             strings.Join(config.SeedTracks, ","),
			"min_acousticness":        strconv.FormatFloat(config.MinAcousticness, 'f', -1, 64),
			"max_acousticness":        strconv.FormatFloat(config.MaxAcousticness, 'f', -1, 64),
			"target_acousticness":     strconv.FormatFloat(config.TargetAcousticness, 'f', -1, 64),
			"min_danceability":        strconv.FormatFloat(config.MinDanceability, 'f', -1, 64),
			"max_danceability":        strconv.FormatFloat(config.MaxDanceability, 'f', -1, 64),
			"target_danceability":     strconv.FormatFloat(config.TargetDanceability, 'f', -1, 64),
			"min_duration_ms":         strconv.Itoa(config.MinDurationMs),
			"max_duration_ms":         strconv.Itoa(config.MaxDurationMs),
			"target_duration_ms":      strconv.Itoa(config.TargetDurationMs),
			"min_energy":              strconv.FormatFloat(config.MinEnergy, 'f', -1, 64),
			"max_energy":              strconv.FormatFloat(config.MaxEnergy, 'f', -1, 64),
			"target_energy":           strconv.FormatFloat(config.TargetEnergy, 'f', -1, 64),
			"min_instrumentalness":    strconv.FormatFloat(config.MinInstrumentalness, 'f', -1, 64),
			"max_instrumentalness":    strconv.FormatFloat(config.MaxInstrumentalness, 'f', -1, 64),
			"target_instrumentalness": strconv.FormatFloat(config.TargetInstrumentalness, 'f', -1, 64),
			"min_key":                 strconv.Itoa(config.MinKey),
			"max_key":                 strconv.Itoa(config.MaxKey),
			"target_key":              strconv.Itoa(config.TargetKey),
			"min_liveness":            strconv.FormatFloat(config.MinLiveness, 'f', -1, 64),
			"max_liveness":            strconv.FormatFloat(config.MaxLiveness, 'f', -1, 64),
			"target_liveness":         strconv.FormatFloat(config.TargetLiveness, 'f', -1, 64),
			"min_loudness":            strconv.FormatFloat(config.MinLoudness, 'f', -1, 64),
			"max_loudness":            strconv.FormatFloat(config.MaxLoudness, 'f', -1, 64),
			"target_loudness":         strconv.FormatFloat(config.TargetLoudness, 'f', -1, 64),
			"min_mode":                strconv.Itoa(config.MinMode),
			"max_mode":                strconv.Itoa(config.MaxMode),
			"target_mode":             strconv.Itoa(config.TargetMode),
			"min_popularity":          strconv.Itoa(config.MinPopularity),
			"max_popularity":          strconv.Itoa(config.MaxPopularity),
			"target_popularity":       strconv.Itoa(config.TargetPopularity),
			"min_speechiness":         strconv.FormatFloat(config.MinSpeechiness, 'f', -1, 64),
			"max_speechiness":         strconv.FormatFloat(config.MaxSpeechiness, 'f', -1, 64),
			"target_speechiness":      strconv.FormatFloat(config.TargetSpeechiness, 'f', -1, 64),
			"min_tempo":               strconv.FormatFloat(config.MinTempo, 'f', -1, 64),
			"max_tempo":               strconv.FormatFloat(config.MaxTempo, 'f', -1, 64),
			"target_tempo":            strconv.FormatFloat(config.TargetTempo, 'f', -1, 64),
			"min_time_signature":      strconv.Itoa(config.MinTimeSignature),
			"max_time_signature":      strconv.Itoa(config.MaxTimeSignature),
			"target_time_signature":   strconv.Itoa(config.TargetTimeSignature),
			"min_valence":             strconv.FormatFloat(config.MinValence, 'f', -1, 64),
			"max_valence":             strconv.FormatFloat(config.MaxValence, 'f', -1, 64),
			"target_valence":          strconv.FormatFloat(config.TargetValence, 'f', -1, 64),
		},
		Headers: map[string]string{
			"Authorization": "Bearer " + config.OauthToken,
		},
		Path: spotifyBaseAddress + recommendationPath,
	})

	if err != nil {
		return nil, &errors.HttpError{StatusCode: 503, Err: err.Error()}
	}

	if resp.StatusCode == http.StatusOK {
		genericResp, err := io.UnmarshalGenericFunction(resp.Body, models.TracksPaging{})
		mapstructure.Decode(genericResp, &recommendationResponse)
		return recommendationResponse, err
	}
	return nil, &errors.HttpError{StatusCode: resp.StatusCode, Err: resp.Status}
}
