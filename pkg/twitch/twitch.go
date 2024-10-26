package twitch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Channel struct {
	ID           string
	UserID       string
	UserLogin    string
	UserName     string
	GameID       string
	GameName     string
	Type         string
	Title        string
	Viewers      int
	StartedAt    time.Time
	Language     string
	ThumbnailURL string
	TagIDs       []string
	Tags         []string
	IsMature     bool
}

type twitchResponse struct {
	Data []struct {
		ID           string    `json:"id"`
		UserID       string    `json:"user_id"`
		UserLogin    string    `json:"user_login"`
		UserName     string    `json:"user_name"`
		GameID       string    `json:"game_id"`
		GameName     string    `json:"game_name"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailURL string    `json:"thumbnail_url"`
		TagIDs       []string  `json:"tag_ids"`
		Tags         []string  `json:"tags"`
		IsMature     bool      `json:"is_mature"`
	} `json:"data"`
}

func getAccessToken(clientID, clientSecret string) (string, error) {
	uri := "https://id.twitch.tv/oauth2/token"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token, status code: %d", resp.StatusCode)
	}

	var result struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}

func GetStreamData(channel, name string) *Channel {
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	clientSecret := os.Getenv("TWITCH_CLIENT_SECRET")

	token, err := getAccessToken(clientID, clientSecret)
	if err != nil {
		fmt.Println("Error obtaining access token:", err)
		return nil
	}

	url := fmt.Sprintf("https://api.twitch.tv/helix/streams?user_login=%s", channel)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Client-ID", clientID)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return nil
	}

	var twitchResp twitchResponse
	if err := json.NewDecoder(resp.Body).Decode(&twitchResp); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	if len(twitchResp.Data) == 0 {
		fmt.Println("El canal no está en directo.")
		return nil
	}

	stream := twitchResp.Data[0]
	return &Channel{
		UserName:     name,
		UserLogin:    stream.UserLogin,
		Viewers:      stream.ViewerCount,
		Title:        stream.Title,
		ID:           stream.ID,
		UserID:       stream.UserID,
		GameID:       stream.GameID,
		GameName:     stream.GameName,
		Type:         stream.Type,
		StartedAt:    stream.StartedAt,
		Language:     stream.Language,
		ThumbnailURL: stream.ThumbnailURL,
		TagIDs:       stream.TagIDs,
		Tags:         stream.Tags,
		IsMature:     stream.IsMature,
	}
}
