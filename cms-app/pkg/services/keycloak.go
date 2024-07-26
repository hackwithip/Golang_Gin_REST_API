package services

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const GrantType string = "password"
const ClientId string = "cms-client"
const ClientSecret string = "zHJsC5y8wRh8eVJosTA5dAJQSCJm2E7P"
const Username string = "admin"
const Password string = "inder@123"
const KeyCloakBaseUrl string = "http://localhost:8080"

type KeyCloakAdminTokenResponse struct {
	AccessToken        string `json:"access_token"`
	ExpiresIn          int    `json:"expires_in"`
	RefreshExpiresIn   int    `json:"refresh_expires_in"`
	RefreshToken       string `json:"refresh_token"`
	TokenType          string `json:"token_type"`
	NotBeforePolicy    int    `json:"not-before-policy"`
	SessionState       string `json:"session_state"`
	Scope              string `json:"scope"`
}


func GetKeycloakAdminToken() (*KeyCloakAdminTokenResponse, error) {
	apiEndpoint := "/realms/master/protocol/openid-connect/token"
	apiURL := KeyCloakBaseUrl + apiEndpoint
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "cms-client")
	data.Set("client_secret", "zHJsC5y8wRh8eVJosTA5dAJQSCJm2E7P")
	data.Set("username", "admin")
	data.Set("password", "inder@123")

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// Set the content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResp KeyCloakAdminTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}
	// Return the response body as a string
	return &tokenResp, nil
}

func CreateKeycloakUser() {

}

func GetKeycloakUser() {

}

func GetKeycloakAuthToken() {

}