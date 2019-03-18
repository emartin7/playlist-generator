package models

type RecommendationRequest struct {
	RecommendationLimit    int        `json:"recommendation_limit"`
	Market                 *string    `json:"market"`
	MinAcousticness        float64    `json:"min_acousticness"`
	MaxAcousticness        float64    `json:"max_acousticness"`
	TargetAcousticness     float64    `json:"target_acousticness"`
	MinDanceability        float64    `json:"min_danceability"`
	MaxDanceability        float64    `json:"max_danceability"`
	TargetDanceability     float64    `json:"target_danceability"`
	MinDurationMs          int        `json:"min_duration_ms"`
	MaxDurationMs          int        `json:"max_duration_ms"`
	TargetDurationMs       int        `json:"target_duration_ms"`
	MinEnergy              float64    `json:"min_energy"`
	MaxEnergy              float64    `json:"max_energy"`
	TargetEnergy           float64    `json:"target_energy"`
	MinInstrumentalness    float64    `json:"min_instrumentalness"`
	MaxInstrumentalness    float64    `json:"max_instrumentalness"`
	TargetInstrumentalness float64    `json:"target_instrumentalness"`
	MinKey                 int        `json:"min_key"`
	MaxKey                 int        `json:"max_key"`
	TargetKey              int        `json:"target_key"`
	MinLiveness            float64    `json:"min_liveness"`
	MaxLiveness            float64    `json:"max_liveness"`
	TargetLiveness         float64    `json:"target_liveness"`
	MinLoudness            float64    `json:"min_loudness"`
	MaxLoudness            float64    `json:"max_loudness"`
	TargetLoudness         float64    `json:"target_loudness"`
	MinMode                int        `json:"min_mode"`
	MaxMode                int        `json:"max_mode"`
	TargetMode             int        `json:"target_mode"`
	MinPopularity          int        `json:"min_popularity"`
	MaxPopularity          int        `json:"max_popularity"`
	TargetPopularity       int        `json:"target_popularity"`
	MinSpeechiness         float64    `json:"min_speechiness"`
	MaxSpeechiness         float64    `json:"max_speechiness"`
	TargetSpeechiness      float64    `json:"target_speechiness"`
	MinTempo               float64    `json:"min_tempo"`
	MaxTempo               float64    `json:"max_tempo"`
	TargetTempo            float64    `json:"target_tempo"`
	MinTimeSignature       int        `json:"min_time_signature"`
	MaxTimeSignature       int        `json:"max_time_signature"`
	TargetTimeSignature    int        `json:"target_time_signature"`
	MinValence             float64    `json:"min_valence"`
	MaxValence             float64    `json:"max_valence"`
	TargetValence          float64    `json:"target_valence"`
	UseUserHistory         bool       `json:"use_user_history"`
	TypeOfSearch           string     `json:"typeOfSearch"`
	TimeRange              string     `json:"timeRange"`
	Offset                 int        `json:"offset"`
	Limit                  int        `json:"limit"`
	OauthToken             string `json:"oauthToken"`
	SeedArtists            []string   `json:"seed_artists"`
	SeedGenres             []string   `json:"seed_genres"`
	SeedTracks             []string   `json:"seed_tracks"`
}

type Seeds struct {
	SeedArtists []string `json:"seed_artists"`
	SeedGenres  []string `json:"seed_genres"`
	SeedTracks  []string `json:"seed_tracks"`
}
