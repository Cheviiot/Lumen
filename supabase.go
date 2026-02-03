package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Supabase handles Supabase API interactions for cloud sync
type Supabase struct {
	client  *http.Client
	baseURL string
	apiKey  string
	userID  string
}

// SyncData represents data to sync with Supabase
type SyncData struct {
	UserID    string     `json:"user_id"`
	Favorites []Favorite `json:"favorites"`
	History   []History  `json:"history"`
	Bookmarks []Bookmark `json:"bookmarks"`
	UpdatedAt string     `json:"updated_at"`
}

// NewSupabase creates a new Supabase instance
func NewSupabase() *Supabase {
	return &Supabase{
		client: &http.Client{},
	}
}

// SetConfig sets Supabase connection settings
func (s *Supabase) SetConfig(baseURL, apiKey, userID string) {
	s.baseURL = baseURL
	s.apiKey = apiKey
	s.userID = userID
}

// request makes an API request to Supabase
func (s *Supabase) request(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, s.baseURL+"/rest/v1"+endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", s.apiKey)
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// SyncFavorites syncs favorites to Supabase
func (s *Supabase) SyncFavorites(favorites []Favorite) error {
	for _, fav := range favorites {
		data := map[string]interface{}{
			"user_id":    s.userID,
			"tmdb_id":    fav.TMDBID,
			"media_type": fav.MediaType,
			"title":      fav.Title,
			"poster":     fav.Poster,
			"year":       fav.Year,
			"rating":     fav.Rating,
			"created_at": fav.CreatedAt,
		}

		_, err := s.request("POST", "/favorites", data)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFavorites retrieves favorites from Supabase
func (s *Supabase) GetFavorites() ([]Favorite, error) {
	data, err := s.request("GET", fmt.Sprintf("/favorites?user_id=eq.%s&order=created_at.desc", s.userID), nil)
	if err != nil {
		return nil, err
	}

	var favorites []Favorite
	if err := json.Unmarshal(data, &favorites); err != nil {
		return nil, err
	}

	return favorites, nil
}

// SyncHistory syncs history to Supabase
func (s *Supabase) SyncHistory(history []History) error {
	for _, h := range history {
		data := map[string]interface{}{
			"user_id":      s.userID,
			"tmdb_id":      h.TMDBID,
			"media_type":   h.MediaType,
			"title":        h.Title,
			"poster":       h.Poster,
			"episode":      h.Episode,
			"progress":     h.Progress,
			"duration":     h.Duration,
			"watched_at":   h.WatchedAt,
			"last_watched": h.LastWatched,
		}

		_, err := s.request("POST", "/history", data)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetHistory retrieves history from Supabase
func (s *Supabase) GetHistory() ([]History, error) {
	data, err := s.request("GET", fmt.Sprintf("/history?user_id=eq.%s&order=last_watched.desc", s.userID), nil)
	if err != nil {
		return nil, err
	}

	var history []History
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return history, nil
}

// SyncBookmarks syncs bookmarks to Supabase
func (s *Supabase) SyncBookmarks(bookmarks []Bookmark) error {
	for _, b := range bookmarks {
		data := map[string]interface{}{
			"user_id":    s.userID,
			"tmdb_id":    b.TMDBID,
			"media_type": b.MediaType,
			"title":      b.Title,
			"poster":     b.Poster,
			"year":       b.Year,
			"note":       b.Note,
			"created_at": b.CreatedAt,
		}

		_, err := s.request("POST", "/bookmarks", data)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetBookmarks retrieves bookmarks from Supabase
func (s *Supabase) GetBookmarks() ([]Bookmark, error) {
	data, err := s.request("GET", fmt.Sprintf("/bookmarks?user_id=eq.%s&order=created_at.desc", s.userID), nil)
	if err != nil {
		return nil, err
	}

	var bookmarks []Bookmark
	if err := json.Unmarshal(data, &bookmarks); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

// TestConnection tests if Supabase is accessible
func (s *Supabase) TestConnection() (bool, error) {
	req, err := http.NewRequest("GET", s.baseURL+"/rest/v1/", nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("apikey", s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// FullSync performs a full sync of all data
func (s *Supabase) FullSync(db *Database) error {
	// Get local data
	favorites, err := db.GetFavorites()
	if err != nil {
		return err
	}

	history, err := db.GetHistory()
	if err != nil {
		return err
	}

	bookmarks, err := db.GetBookmarks()
	if err != nil {
		return err
	}

	// Sync to cloud
	if err := s.SyncFavorites(favorites); err != nil {
		return err
	}

	if err := s.SyncHistory(history); err != nil {
		return err
	}

	if err := s.SyncBookmarks(bookmarks); err != nil {
		return err
	}

	return nil
}
