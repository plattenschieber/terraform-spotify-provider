package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	token string
	id    string
}

// Return a new client
func NewClient(token, id string) *Client {

	return &Client{token, id}
}

const BASE_URL = "https://api.spotify.com/v1/"

type CreatePlayListRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

type AddTracksToPlayListRequest struct {
	Tracks []string `json:"uris"`
}
type Track struct {
	URI string `json:uri`
}
type RemoveTracksFromPlayListRequest struct {
	Tracks []Track `json:"tracks"`
}

// https://api.spotify.com/v1/users/{user_id}/playlists
//
//	{
//	 "name": "New Playlist from curl2",
//	 "description": "New playlist description",
//	 "public": false
//	}
func (c *Client) CreatePlayList(name, description string, public bool) (string, error) {
	userId := c.id
	playlistRequest := CreatePlayListRequest{
		Name:        name,
		Description: description,
		Public:      public,
	}

	d, err := json.Marshal(playlistRequest)
	if err != nil {
		return "", fmt.Errorf("an error happend during marshalling data: %s", err.Error())
	}

	resp, err := http.DefaultClient.Post(fmt.Sprintf("%susers/%s/playlists", BASE_URL, userId), "application/json", bytes.NewReader(d))

	if err != nil {
		return "", fmt.Errorf("an error happend during CreatePlayList(): %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("got status code %d, expected 200", resp.StatusCode)
	}

	output := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return "", fmt.Errorf("unable to decode response: %s", err)
	}

	return output["id"].(string), nil
}

// https://api.spotify.com/v1//playlists/{playlist_id}/followers

func (c *Client) UnfollowPlayList(playlistId string) (string, error) {

	resp, err := http.DefaultClient.Post(fmt.Sprintf("%splaylists/%s/followers", BASE_URL, playlistId), "application/json", bytes.NewReader(d))

	if err != nil {
		return "", fmt.Errorf("an error happend during UnfollowPlayList(): %s", err.Error())
	}

	defer resp.Body.Close()

	if resp != nil {
		return "", fmt.Errorf("got status code %d, expected 200", resp.StatusCode)
	}

	return ("Successfully unfollowed " + playlistId), nil
}

// https://api.spotify.com/v1/playlists/playlist_id/tracks
//
//	{
//	 "uris": [spotify:track:123,spotify:track:456]
//	}

func (c *Client) AddTracksToPlayList(playlistId string, tracks []string) (string, error) {

	playlistRequest := AddTracksToPlayListRequest{
		Tracks: tracks,
	}

	d, err := json.Marshal(playlistRequest)
	if err != nil {
		return "", fmt.Errorf("an error happend during marshalling data: %s", err.Error())
	}

	resp, err := http.DefaultClient.Post(fmt.Sprintf("%splaylists/%s/tracks", BASE_URL, playlistId), "application/json", bytes.NewReader(d))

	if err != nil {
		return "", fmt.Errorf("an error happend during UnfollowPlayList(): %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("got status code %d, expected 200", resp.StatusCode)
	}

	output := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return "", fmt.Errorf("unable to decode response: %s", err)
	}

	return output["snapshot_id"].(string), nil
}

// https://api.spotify.com/v1/playlists/{playlist_id}/tracks
//
//	{
//	 "tracks": [{ "uri": "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" },{ "uri": "spotify:track:1301WleyT98MSxVHPZCA6M" }]
//	}

func (c *Client) RemoveTracksFromPlayList(playlistId string, uris []string) (string, error) {

	var tracks []Track
	for i, s := range uris {
		t := Track{
			URI: s,
		}
		tracks = append(tracks, t)
	}

	playlistRequest := RemoveTracksFromPlayListRequest{
		Tracks: tracks,
	}

	d, err := json.Marshal(playlistRequest)
	if err != nil {
		return "", fmt.Errorf("an error happend during marshalling data: %s", err.Error())
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%splaylists/%s/tracks", BASE_URL, playlistId), bytes.NewReader(d))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("an error happend during UnfollowPlayList(): %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("got status code %d, expected 200", resp.StatusCode)
	}

	output := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return "", fmt.Errorf("unable to decode response: %s", err)
	}

	return output["snapshot_id"].(string), nil
}

//TODO:: Get Track
// https://api.spotify.com/v1/search?q=Pink%20Venom%20Blackpink&type=track&limit=5
// q = Keywords
// type = track
// limit = 5
// response
// artists[0][name]
// name
// uri
