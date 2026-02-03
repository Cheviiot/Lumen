package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Database handles local JSON storage
type Database struct {
	dataDir string
}

// Favorite model for storing favorite items
type Favorite struct {
	ID        uint      `json:"id"`
	TMDBID    int       `json:"tmdb_id"`
	MediaType string    `json:"media_type"`
	Title     string    `json:"title"`
	Poster    string    `json:"poster"`
	Year      string    `json:"year"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

// History model for storing watch history
type History struct {
	ID          uint      `json:"id"`
	TMDBID      int       `json:"tmdb_id"`
	MediaType   string    `json:"media_type"`
	Title       string    `json:"title"`
	Poster      string    `json:"poster"`
	Episode     string    `json:"episode,omitempty"`
	Progress    float64   `json:"progress"`
	Duration    int       `json:"duration"`
	WatchedAt   time.Time `json:"watched_at"`
	LastWatched time.Time `json:"last_watched"`
}

// Bookmark model for storing bookmarks
type Bookmark struct {
	ID        uint      `json:"id"`
	TMDBID    int       `json:"tmdb_id"`
	MediaType string    `json:"media_type"`
	Title     string    `json:"title"`
	Poster    string    `json:"poster"`
	Year      string    `json:"year"`
	Note      string    `json:"note,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// Settings model for app settings
type Settings struct {
	ID                  uint   `gorm:"primaryKey"`
	TorrServerURL       string `json:"torr_server_url"`
	JackettURL          string `json:"jackett_url"`
	JackettAPIKey       string `json:"jackett_api_key"`
	JackettPublicParser string `json:"jackett_public_parser"`
	SupabaseURL         string `json:"supabase_url"`
	SupabaseKey         string `json:"supabase_key"`
	Theme               string `json:"theme"`
	Language            string `json:"language"`
	AutoSync            bool   `json:"auto_sync"`
	DownloadPath        string `json:"download_path"`
	PlayerPath          string `json:"player_path"`
}

// Storage структуры для JSON файлов
type FavoritesStorage struct {
	Items      []Favorite `json:"items"`
	LastID     uint       `json:"last_id"`
	LastUpdate time.Time  `json:"last_update"`
}

type HistoryStorage struct {
	Items      []History `json:"items"`
	LastID     uint      `json:"last_id"`
	LastUpdate time.Time `json:"last_update"`
}

type BookmarksStorage struct {
	Items      []Bookmark `json:"items"`
	LastID     uint       `json:"last_id"`
	LastUpdate time.Time  `json:"last_update"`
}

// NewDatabase creates a new Database instance
func NewDatabase() *Database {
	return &Database{}
}

// Init initializes the database (creates data directory)
func (d *Database) Init() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.Getenv("HOME")
	}

	d.dataDir = filepath.Join(configDir, "Lumen")
	return os.MkdirAll(d.dataDir, 0755)
}

// Helper functions for file operations
func (d *Database) readJSON(filename string, v interface{}) error {
	path := filepath.Join(d.dataDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Файл не существует - не ошибка
		}
		return err
	}
	return json.Unmarshal(data, v)
}

func (d *Database) writeJSON(filename string, v interface{}) error {
	path := filepath.Join(d.dataDir, filename)
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// GetFavorites returns all favorites
func (d *Database) GetFavorites() ([]Favorite, error) {
	var storage FavoritesStorage
	if err := d.readJSON("favorites.json", &storage); err != nil {
		return nil, err
	}
	if storage.Items == nil {
		storage.Items = []Favorite{}
	}
	return storage.Items, nil
}

// AddFavorite adds a new favorite
func (d *Database) AddFavorite(favorite Favorite) error {
	var storage FavoritesStorage
	d.readJSON("favorites.json", &storage)

	if storage.Items == nil {
		storage.Items = []Favorite{}
	}

	storage.LastID++
	favorite.ID = storage.LastID
	favorite.CreatedAt = time.Now()
	storage.Items = append(storage.Items, favorite)
	storage.LastUpdate = time.Now()

	return d.writeJSON("favorites.json", &storage)
}

// RemoveFavorite removes a favorite by TMDB ID
func (d *Database) RemoveFavorite(tmdbID int, mediaType string) error {
	var storage FavoritesStorage
	if err := d.readJSON("favorites.json", &storage); err != nil {
		return err
	}

	filtered := []Favorite{}
	for _, fav := range storage.Items {
		if !(fav.TMDBID == tmdbID && fav.MediaType == mediaType) {
			filtered = append(filtered, fav)
		}
	}

	storage.Items = filtered
	storage.LastUpdate = time.Now()
	return d.writeJSON("favorites.json", &storage)
}

// IsFavorite checks if item is in favorites
func (d *Database) IsFavorite(tmdbID int, mediaType string) bool {
	var storage FavoritesStorage
	if err := d.readJSON("favorites.json", &storage); err != nil {
		return false
	}

	for _, fav := range storage.Items {
		if fav.TMDBID == tmdbID && fav.MediaType == mediaType {
			return true
		}
	}
	return false
}

// GetHistory returns watch history
func (d *Database) GetHistory() ([]History, error) {
	var storage HistoryStorage
	if err := d.readJSON("history.json", &storage); err != nil {
		return nil, err
	}
	if storage.Items == nil {
		storage.Items = []History{}
	}
	return storage.Items, nil
}

// AddHistory adds or updates history entry
func (d *Database) AddHistory(history History) error {
	var storage HistoryStorage
	d.readJSON("history.json", &storage)

	if storage.Items == nil {
		storage.Items = []History{}
	}

	// Ищем существующую запись
	found := false
	for i := range storage.Items {
		if storage.Items[i].TMDBID == history.TMDBID && storage.Items[i].MediaType == history.MediaType {
			storage.Items[i].Progress = history.Progress
			storage.Items[i].LastWatched = time.Now()
			storage.Items[i].Episode = history.Episode
			found = true
			break
		}
	}

	if !found {
		storage.LastID++
		history.ID = storage.LastID
		history.WatchedAt = time.Now()
		history.LastWatched = time.Now()
		storage.Items = append(storage.Items, history)
	}

	storage.LastUpdate = time.Now()
	return d.writeJSON("history.json", &storage)
}

// ClearHistory clears all history
func (d *Database) ClearHistory() error {
	storage := HistoryStorage{
		Items:      []History{},
		LastID:     0,
		LastUpdate: time.Now(),
	}
	return d.writeJSON("history.json", &storage)
}

// GetBookmarks returns all bookmarks
func (d *Database) GetBookmarks() ([]Bookmark, error) {
	var storage BookmarksStorage
	if err := d.readJSON("bookmarks.json", &storage); err != nil {
		return nil, err
	}
	if storage.Items == nil {
		storage.Items = []Bookmark{}
	}
	return storage.Items, nil
}

// AddBookmark adds a new bookmark
func (d *Database) AddBookmark(bookmark Bookmark) error {
	var storage BookmarksStorage
	d.readJSON("bookmarks.json", &storage)

	if storage.Items == nil {
		storage.Items = []Bookmark{}
	}

	storage.LastID++
	bookmark.ID = storage.LastID
	bookmark.CreatedAt = time.Now()
	storage.Items = append(storage.Items, bookmark)
	storage.LastUpdate = time.Now()

	return d.writeJSON("bookmarks.json", &storage)
}

// RemoveBookmark removes a bookmark
func (d *Database) RemoveBookmark(tmdbID int, mediaType string) error {
	var storage BookmarksStorage
	if err := d.readJSON("bookmarks.json", &storage); err != nil {
		return err
	}

	filtered := []Bookmark{}
	for _, bm := range storage.Items {
		if !(bm.TMDBID == tmdbID && bm.MediaType == mediaType) {
			filtered = append(filtered, bm)
		}
	}

	storage.Items = filtered
	storage.LastUpdate = time.Now()
	return d.writeJSON("bookmarks.json", &storage)
}

// IsBookmarked checks if item is bookmarked
func (d *Database) IsBookmarked(tmdbID int, mediaType string) bool {
	var storage BookmarksStorage
	if err := d.readJSON("bookmarks.json", &storage); err != nil {
		return false
	}

	for _, bm := range storage.Items {
		if bm.TMDBID == tmdbID && bm.MediaType == mediaType {
			return true
		}
	}
	return false
}

// GetSettings returns app settings
func (d *Database) GetSettings() (Settings, error) {
	var settings Settings
	if err := d.readJSON("settings.json", &settings); err != nil {
		// Возвращаем настройки по умолчанию
		return Settings{
			ID:            1,
			TorrServerURL: "http://localhost:8090",
			JackettURL:    "http://localhost:9117",
			Theme:         "dark",
			Language:      "ru",
			AutoSync:      false,
		}, nil
	}

	// Если файл пустой, вернуть defaults
	if settings.ID == 0 {
		return Settings{
			ID:            1,
			TorrServerURL: "http://localhost:8090",
			JackettURL:    "http://localhost:9117",
			Theme:         "dark",
			Language:      "ru",
			AutoSync:      false,
		}, nil
	}

	return settings, nil
}

// SaveSettings saves app settings
func (d *Database) SaveSettings(settings Settings) error {
	if settings.ID == 0 {
		settings.ID = 1
	}
	return d.writeJSON("settings.json", &settings)
}
