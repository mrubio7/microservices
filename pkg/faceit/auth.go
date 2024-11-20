package faceit

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

type userInfo struct {
	GUID     string `json:"guid"`
	Nickname string `json:"nickname"`
}

func Auth(code, codeVerifier string) (*model.UserModel, error) {
	if code == "" || codeVerifier == "" {
		return nil, errors.New("code or code_verifier doesn't exist")
	}

	clientID := os.Getenv("FACEIT_CLIENT_ID")
	clientSecret := os.Getenv("FACEIT_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return nil, errors.New("server configuration error")
	}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("code_verifier", codeVerifier)

	logger.Info("%s - %s", clientID, clientSecret)
	credentials := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	logger.Info("%s", credentials)

	req, err := http.NewRequest("POST", "https://api.faceit.com/auth/v1/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, errors.New("error creating http request")
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+credentials)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("http request error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error getting tokens")
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, errors.New("error decoding body response")
	}

	user, err := getUserInfo(tokenResponse.AccessToken)
	if err != nil {
		return nil, errors.New("error getting user info")
	}

	return user, nil
}

func getUserInfo(accessToken string) (*model.UserModel, error) {
	req, err := http.NewRequest("GET", "https://api.faceit.com/auth/v1/resources/userinfo", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al obtener información del usuario: %s", string(body))
	}

	var userInfo userInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	user := &model.UserModel{
		FaceitId: userInfo.GUID,
		Name:     userInfo.Nickname,
	}

	return user, nil
}
