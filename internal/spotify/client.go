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
