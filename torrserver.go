package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TorrServer handles TorrServer API interactions
type TorrServer struct {
	client  *http.Client
	baseURL string
}

// Torrent represents a torrent in TorrServer
type Torrent struct {
	Title     string `json:"title"`
	Hash      string `json:"hash"`
	Poster    string `json:"poster"`
	Data      string `json:"data"`
	Timestamp int64  `json:"timestamp"`
	Size      int64  `json:"size"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	Stat      int    `json:"stat"`
}

// TorrentFile represents a file within a torrent
type TorrentFile struct {
	ID     int    `json:"id"`
	Path   string `json:"path"`
	Length int64  `json:"length"`
}

// TorrentStat represents torrent statistics
type TorrentStat struct {
	Hash             string        `json:"hash"`
	Title            string        `json:"title"`
	TorrentSize      int64         `json:"torrent_size"`
	DownloadSpeed    float64       `json:"download_speed"`
	UploadSpeed      float64       `json:"upload_speed"`
	ActivePeers      int           `json:"active_peers"`
	TotalPeers       int           `json:"total_peers"`
	ConnectedSeeders int           `json:"connected_seeders"`
	FileStats        []TorrentFile `json:"file_stats"`
}

// NewTorrServer creates a new TorrServer instance
func NewTorrServer() *TorrServer {
	return &TorrServer{
		client:  &http.Client{},
		baseURL: "http://localhost:8090",
	}
}

// SetBaseURL sets the TorrServer base URL
func (t *TorrServer) SetBaseURL(url string) {
	t.baseURL = url
}

// request makes an API request to TorrServer
func (t *TorrServer) request(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, t.baseURL+endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("TorrServer response status: %d\n", resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("TorrServer returned error: %d - %s", resp.StatusCode, string(data))
	}

	return data, nil
}

// Echo checks if TorrServer is running
func (t *TorrServer) Echo() (bool, error) {
	resp, err := t.client.Get(t.baseURL + "/echo")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}

// AddTorrent adds a torrent by magnet link
func (t *TorrServer) AddTorrent(magnetLink, title, poster string, saveToDb bool) (*Torrent, error) {
	body := map[string]interface{}{
		"action":     "add",
		"link":       magnetLink,
		"title":      title,
		"poster":     poster,
		"save_to_db": saveToDb,
	}

	jsonBody, _ := json.Marshal(body)
	fmt.Printf("TorrServer request: %s\n", string(jsonBody))

	data, err := t.request("POST", "/torrents", body)
	if err != nil {
		fmt.Printf("Request error: %v\n", err)
		return nil, err
	}

	fmt.Printf("TorrServer response: %s\n", string(data))

	if len(data) == 0 {
		return nil, fmt.Errorf("empty response from TorrServer")
	}

	var torrent Torrent
	if err := json.Unmarshal(data, &torrent); err != nil {
		fmt.Printf("Failed to unmarshal response: %v\n", err)
		return nil, err
	}

	fmt.Printf("Successfully added torrent with hash: %s\n", torrent.Hash)

	return &torrent, nil
}

// AddTorrentFile adds a torrent by file content (base64)
func (t *TorrServer) AddTorrentFile(fileContent, title, poster string, saveToDb bool) (*Torrent, error) {
	body := map[string]interface{}{
		"action":     "add",
		"link":       fileContent,
		"title":      title,
		"poster":     poster,
		"save_to_db": saveToDb,
	}

	data, err := t.request("POST", "/torrents", body)
	if err != nil {
		return nil, err
	}

	var torrent Torrent
	if err := json.Unmarshal(data, &torrent); err != nil {
		return nil, err
	}

	return &torrent, nil
}

// GetTorrents returns list of all torrents
func (t *TorrServer) GetTorrents() ([]Torrent, error) {
	body := map[string]interface{}{
		"action": "list",
	}

	data, err := t.request("POST", "/torrents", body)
	if err != nil {
		return nil, err
	}

	var torrents []Torrent
	if err := json.Unmarshal(data, &torrents); err != nil {
		return nil, err
	}

	return torrents, nil
}

// GetTorrent returns a specific torrent by hash
func (t *TorrServer) GetTorrent(hash string) (*Torrent, error) {
	body := map[string]interface{}{
		"action": "get",
		"hash":   hash,
	}

	data, err := t.request("POST", "/torrents", body)
	if err != nil {
		return nil, err
	}

	var torrent Torrent
	if err := json.Unmarshal(data, &torrent); err != nil {
		return nil, err
	}

	return &torrent, nil
}

// RemoveTorrent removes a torrent by hash
func (t *TorrServer) RemoveTorrent(hash string) error {
	body := map[string]interface{}{
		"action": "rem",
		"hash":   hash,
	}

	_, err := t.request("POST", "/torrents", body)
	return err
}

// GetTorrentStat returns torrent statistics
func (t *TorrServer) GetTorrentStat(hash string) (*TorrentStat, error) {
	body := map[string]interface{}{
		"action": "get",
		"hash":   hash,
	}

	data, err := t.request("POST", "/torrents", body)
	if err != nil {
		return nil, err
	}

	var stat TorrentStat
	if err := json.Unmarshal(data, &stat); err != nil {
		return nil, err
	}

	return &stat, nil
}

// GetStreamURL returns the stream URL for a file
func (t *TorrServer) GetStreamURL(hash string, fileIndex int) string {
	url := fmt.Sprintf("%s/stream/%s?index=%d&play", t.baseURL, hash, fileIndex)
	fmt.Printf("GetStreamURL: hash=%s, index=%d, url=%s\n", hash, fileIndex, url)
	return url
}

// PreloadTorrent starts preloading a torrent
func (t *TorrServer) PreloadTorrent(hash string, fileIndex int) error {
	body := map[string]interface{}{
		"action": "preload",
		"hash":   hash,
		"index":  fileIndex,
	}

	_, err := t.request("POST", "/torrents", body)
	return err
}

// DropTorrent drops torrent from memory (keeps in DB if saved)
func (t *TorrServer) DropTorrent(hash string) error {
	body := map[string]interface{}{
		"action": "drop",
		"hash":   hash,
	}

	_, err := t.request("POST", "/torrents", body)
	return err
}

// GetSettings returns TorrServer settings
func (t *TorrServer) GetSettings() (map[string]interface{}, error) {
	data, err := t.request("POST", "/settings", map[string]interface{}{
		"action": "get",
	})
	if err != nil {
		return nil, err
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	return settings, nil
}

// SetSettings updates TorrServer settings
func (t *TorrServer) SetSettings(settings map[string]interface{}) error {
	settings["action"] = "set"
	_, err := t.request("POST", "/settings", settings)
	return err
}
