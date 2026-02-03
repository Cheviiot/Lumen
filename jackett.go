package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Jackett handles Jackett API interactions
type Jackett struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

// JackettResult represents a search result from Jackett
type JackettResult struct {
	Title                string  `json:"Title"`
	Tracker              string  `json:"Tracker"`
	TrackerId            string  `json:"TrackerId"`
	CategoryDesc         string  `json:"CategoryDesc"`
	PublishDate          string  `json:"PublishDate"`
	Size                 int64   `json:"Size"`
	Seeders              int     `json:"Seeders"`
	Peers                int     `json:"Peers"`
	Grabs                int     `json:"Grabs"`
	MagnetUri            string  `json:"MagnetUri"`
	Link                 string  `json:"Link"`
	Guid                 string  `json:"Guid"`
	Details              string  `json:"Details"`
	InfoHash             string  `json:"InfoHash"`
	DownloadVolumeFactor float64 `json:"DownloadVolumeFactor"`
	UploadVolumeFactor   float64 `json:"UploadVolumeFactor"`
}

// JackettResponse represents Jackett API response
type JackettResponse struct {
	Results  []JackettResult  `json:"Results"`
	Indexers []JackettIndexer `json:"Indexers"`
}

// JackettIndexer represents a Jackett indexer
type JackettIndexer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
	Results int    `json:"results"`
	Error   string `json:"error"`
}

// IndexerConfig represents indexer configuration
type IndexerConfig struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Configured bool   `json:"configured"`
	Language   string `json:"language"`
}

// NewJackett creates a new Jackett instance
func NewJackett() *Jackett {
	return &Jackett{
		client:  &http.Client{},
		baseURL: "http://localhost:9117",
		apiKey:  "",
	}
}

// SetConfig sets Jackett connection settings
func (j *Jackett) SetConfig(baseURL, apiKey string) {
	j.baseURL = baseURL
	j.apiKey = apiKey
}

// Search searches for torrents across all indexers
func (j *Jackett) Search(query string, categories []int) (*JackettResponse, error) {
	u, err := url.Parse(j.baseURL + "/api/v2.0/indexers/all/results")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("apikey", j.apiKey)
	q.Set("Query", query)

	if len(categories) > 0 {
		for _, cat := range categories {
			q.Add("Category[]", fmt.Sprintf("%d", cat))
		}
	}

	u.RawQuery = q.Encode()

	resp, err := j.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result JackettResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SearchMovies searches for movies (category 2000)
func (j *Jackett) SearchMovies(query string) (*JackettResponse, error) {
	return j.Search(query, []int{2000})
}

// SearchTV searches for TV shows (category 5000)
func (j *Jackett) SearchTV(query string) (*JackettResponse, error) {
	return j.Search(query, []int{5000})
}

// SearchAnime searches for anime (category 5070)
func (j *Jackett) SearchAnime(query string) (*JackettResponse, error) {
	return j.Search(query, []int{5070})
}

// SearchByIndexer searches using a specific indexer
func (j *Jackett) SearchByIndexer(indexerID, query string, categories []int) (*JackettResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s/api/v2.0/indexers/%s/results", j.baseURL, indexerID))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("apikey", j.apiKey)
	q.Set("Query", query)

	if len(categories) > 0 {
		for _, cat := range categories {
			q.Add("Category[]", fmt.Sprintf("%d", cat))
		}
	}

	u.RawQuery = q.Encode()

	resp, err := j.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result JackettResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetIndexers returns list of configured indexers
func (j *Jackett) GetIndexers() ([]IndexerConfig, error) {
	u, err := url.Parse(j.baseURL + "/api/v2.0/indexers")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("apikey", j.apiKey)
	q.Set("configured", "true")
	u.RawQuery = q.Encode()

	resp, err := j.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var indexers []IndexerConfig
	if err := json.Unmarshal(data, &indexers); err != nil {
		return nil, err
	}

	return indexers, nil
}

// TestConnection tests if Jackett is accessible
func (j *Jackett) TestConnection() (bool, error) {
	u, err := url.Parse(j.baseURL + "/api/v2.0/server/config")
	if err != nil {
		return false, err
	}

	q := u.Query()
	q.Set("apikey", j.apiKey)
	u.RawQuery = q.Encode()

	resp, err := j.client.Get(u.String())
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// FormatSize formats size in bytes to human readable format
func FormatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
