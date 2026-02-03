package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	tmdbAPIKey        = "4ef0d7355d9ffb5151e987764708ce96"
	tmdbBaseURL       = "https://api.themoviedb.org/3"
	tmdbImageURL      = "https://image.tmdb.org/t/p"
	tmdbProxyAPI      = "https://tmdbapi.rootu.top/3"
	tmdbProxyImageURL = "https://tmdbimg.rootu.top/t/p"
)

// TMDB handles TMDB API interactions
type TMDB struct {
	client        *http.Client
	useProxyAPI   bool
	useProxyImage bool
}

// Movie represents a movie from TMDB
type Movie struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	Overview      string  `json:"overview"`
	PosterPath    string  `json:"poster_path"`
	BackdropPath  string  `json:"backdrop_path"`
	ReleaseDate   string  `json:"release_date"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
	Popularity    float64 `json:"popularity"`
	GenreIDs      []int   `json:"genre_ids"`
	Adult         bool    `json:"adult"`
}

// TVShow represents a TV show from TMDB
type TVShow struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	FirstAirDate string  `json:"first_air_date"`
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
	Popularity   float64 `json:"popularity"`
	GenreIDs     []int   `json:"genre_ids"`
}

// Collection represents a movie collection
type Collection struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

// CollectionDetails represents detailed collection info
type CollectionDetails struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	Parts        []Movie `json:"parts"`
}

// MovieDetails represents detailed movie info
type MovieDetails struct {
	Movie
	Runtime             int         `json:"runtime"`
	Genres              []Genre     `json:"genres"`
	Tagline             string      `json:"tagline"`
	Status              string      `json:"status"`
	Budget              int64       `json:"budget"`
	Revenue             int64       `json:"revenue"`
	ImdbID              string      `json:"imdb_id"`
	Homepage            string      `json:"homepage"`
	ProductionCompanies []Company   `json:"production_companies"`
	BelongsToCollection *Collection `json:"belongs_to_collection"`
}

// TVShowDetails represents detailed TV show info
type TVShowDetails struct {
	TVShow
	NumberOfSeasons     int       `json:"number_of_seasons"`
	NumberOfEpisodes    int       `json:"number_of_episodes"`
	Genres              []Genre   `json:"genres"`
	Tagline             string    `json:"tagline"`
	Status              string    `json:"status"`
	EpisodeRunTime      []int     `json:"episode_run_time"`
	Seasons             []Season  `json:"seasons"`
	Homepage            string    `json:"homepage"`
	ProductionCompanies []Company `json:"production_companies"`
}

// Season represents a TV season
type Season struct {
	ID           int    `json:"id"`
	SeasonNumber int    `json:"season_number"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
}

// Episode represents a TV episode
type Episode struct {
	ID            int     `json:"id"`
	EpisodeNumber int     `json:"episode_number"`
	SeasonNumber  int     `json:"season_number"`
	Name          string  `json:"name"`
	Overview      string  `json:"overview"`
	StillPath     string  `json:"still_path"`
	AirDate       string  `json:"air_date"`
	VoteAverage   float64 `json:"vote_average"`
	Runtime       int     `json:"runtime"`
}

// Genre represents a genre
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Company represents a production company
type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LogoPath string `json:"logo_path"`
	Country  string `json:"origin_country"`
}

// Video represents a video (trailer, teaser, etc.)
type Video struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Site      string `json:"site"`
	Size      int    `json:"size"`
	Type      string `json:"type"`
	Official  bool   `json:"official"`
	Published string `json:"published_at"`
}

// VideosResult represents video results
type VideosResult struct {
	ID      int     `json:"id"`
	Results []Video `json:"results"`
}

// SearchResult represents search results
type SearchResult struct {
	Page         int     `json:"page"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Results      []Movie `json:"results"`
}

// TVSearchResult represents TV search results
type TVSearchResult struct {
	Page         int      `json:"page"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
	Results      []TVShow `json:"results"`
}

// SeasonDetails represents season details with episodes
type SeasonDetails struct {
	ID           int       `json:"id"`
	SeasonNumber int       `json:"season_number"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	PosterPath   string    `json:"poster_path"`
	AirDate      string    `json:"air_date"`
	Episodes     []Episode `json:"episodes"`
}

// NewTMDB creates a new TMDB instance
func NewTMDB() *TMDB {
	// Создаем HTTP клиент с настроенным Dialer для предпочтения IPv4
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &TMDB{
		client: &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		},
		useProxyAPI:   true, // По умолчанию включаем прокси
		useProxyImage: true, // По умолчанию включаем прокси для картинок
	}
}

// SetProxySettings updates proxy settings
func (t *TMDB) SetProxySettings(useAPI, useImage bool) {
	t.useProxyAPI = useAPI
	t.useProxyImage = useImage
}

// request makes an API request to TMDB
func (t *TMDB) request(endpoint string, params map[string]string) ([]byte, error) {
	baseURL := tmdbBaseURL
	if t.useProxyAPI {
		baseURL = tmdbProxyAPI
	}

	u, err := url.Parse(baseURL + endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("api_key", tmdbAPIKey)
	q.Set("language", "ru-RU")
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	resp, err := t.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// UpdateProxySettings updates TMDB proxy settings from frontend
func (t *TMDB) UpdateProxySettings(useAPI, useImage bool) {
	t.useProxyAPI = useAPI
	t.useProxyImage = useImage
}

// GetImageURL returns the correct image URL based on proxy settings
func (t *TMDB) GetImageURL(path, size string) string {
	if path == "" {
		return ""
	}
	baseURL := tmdbImageURL
	if t.useProxyImage {
		baseURL = tmdbProxyImageURL
	}
	return fmt.Sprintf("%s/%s%s", baseURL, size, path)
}

// SearchMovies searches for movies
func (t *TMDB) SearchMovies(query string, page int) (*SearchResult, error) {
	data, err := t.request("/search/movie", map[string]string{
		"query": query,
		"page":  fmt.Sprintf("%d", page),
	})
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SearchTV searches for TV shows
func (t *TMDB) SearchTV(query string, page int) (*TVSearchResult, error) {
	data, err := t.request("/search/tv", map[string]string{
		"query": query,
		"page":  fmt.Sprintf("%d", page),
	})
	if err != nil {
		return nil, err
	}

	var result TVSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTrendingMovies gets trending movies
func (t *TMDB) GetTrendingMovies(timeWindow string, page int) (*SearchResult, error) {
	endpoint := fmt.Sprintf("/trending/movie/%s", timeWindow)
	data, err := t.request(endpoint, map[string]string{
		"page": fmt.Sprintf("%d", page),
	})
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTrendingTV gets trending TV shows
func (t *TMDB) GetTrendingTV(timeWindow string, page int) (*TVSearchResult, error) {
	endpoint := fmt.Sprintf("/trending/tv/%s", timeWindow)
	data, err := t.request(endpoint, map[string]string{
		"page": fmt.Sprintf("%d", page),
	})
	if err != nil {
		return nil, err
	}

	var result TVSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetPopularMovies gets popular movies
func (t *TMDB) GetPopularMovies(page int) (*SearchResult, error) {
	data, err := t.request("/movie/popular", map[string]string{
		"page": fmt.Sprintf("%d", page),
	})
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetPopularTV gets popular TV shows
func (t *TMDB) GetPopularTV(page int) (*TVSearchResult, error) {
	data, err := t.request("/tv/popular", map[string]string{
		"page": fmt.Sprintf("%d", page),
	})
	if err != nil {
		return nil, err
	}

	var result TVSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMovieDetails gets detailed movie info
func (t *TMDB) GetMovieDetails(movieID int) (*MovieDetails, error) {
	endpoint := fmt.Sprintf("/movie/%d", movieID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result MovieDetails
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTVDetails gets detailed TV show info
func (t *TMDB) GetTVDetails(tvID int) (*TVShowDetails, error) {
	endpoint := fmt.Sprintf("/tv/%d", tvID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result TVShowDetails
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetSeasonDetails gets season details with episodes
func (t *TMDB) GetSeasonDetails(tvID int, seasonNumber int) (*SeasonDetails, error) {
	endpoint := fmt.Sprintf("/tv/%d/season/%d", tvID, seasonNumber)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result SeasonDetails
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAnimationMovies gets animation movies
func (t *TMDB) GetAnimationMovies(page int) (*SearchResult, error) {
	data, err := t.request("/discover/movie", map[string]string{
		"page":        fmt.Sprintf("%d", page),
		"with_genres": "16", // Animation genre ID
		"sort_by":     "popularity.desc",
	})
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAnimationTV gets animation TV shows
func (t *TMDB) GetAnimationTV(page int) (*TVSearchResult, error) {
	data, err := t.request("/discover/tv", map[string]string{
		"page":        fmt.Sprintf("%d", page),
		"with_genres": "16", // Animation genre ID
		"sort_by":     "popularity.desc",
	})
	if err != nil {
		return nil, err
	}

	var result TVSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGenres returns movie genres
func (t *TMDB) GetGenres() ([]Genre, error) {
	data, err := t.request("/genre/movie/list", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Genres []Genre `json:"genres"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result.Genres, nil
}

// GetTVGenres returns TV genres
func (t *TMDB) GetTVGenres() ([]Genre, error) {
	data, err := t.request("/genre/tv/list", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Genres []Genre `json:"genres"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result.Genres, nil
}

// GetMovieVideos returns videos (trailers, teasers) for a movie
func (t *TMDB) GetMovieVideos(movieID int) (*VideosResult, error) {
	endpoint := fmt.Sprintf("/movie/%d/videos", movieID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result VideosResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTVVideos returns videos (trailers, teasers) for a TV show
func (t *TMDB) GetTVVideos(tvID int) (*VideosResult, error) {
	endpoint := fmt.Sprintf("/tv/%d/videos", tvID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result VideosResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMovieSimilar returns similar movies
func (t *TMDB) GetMovieSimilar(movieID int) (*SearchResult, error) {
	endpoint := fmt.Sprintf("/movie/%d/similar", movieID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetMovieRecommendations returns recommended movies
func (t *TMDB) GetMovieRecommendations(movieID int) (*SearchResult, error) {
	endpoint := fmt.Sprintf("/movie/%d/recommendations", movieID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTVSimilar returns similar TV shows
func (t *TMDB) GetTVSimilar(tvID int) (*TVSearchResult, error) {
	endpoint := fmt.Sprintf("/tv/%d/similar", tvID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result TVSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTVRecommendations returns recommended TV shows
func (t *TMDB) GetTVRecommendations(tvID int) (*TVSearchResult, error) {
	endpoint := fmt.Sprintf("/tv/%d/recommendations", tvID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result TVSearchResult
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCollection returns collection details
func (t *TMDB) GetCollection(collectionID int) (*CollectionDetails, error) {
	endpoint := fmt.Sprintf("/collection/%d", collectionID)
	data, err := t.request(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result CollectionDetails
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
