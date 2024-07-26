package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/inder231/cms-app/pkg/utils"
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

type Credentials struct {
	Type string `json:"type"`
	Value string `json:"value"`
	Temporary bool `json:"temporary"`
}

type KeycloakUser struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Enabled bool `json:"enabled"`
	Credentials []Credentials `json:"credentials"`
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

func CreateKeycloakUser(adminToken, email string) (string, bool) {
	apiEndpoint := "/admin/realms/master/users"
	apiURL := KeyCloakBaseUrl + apiEndpoint

	// Generate random password to store in keycloak
	randomPassword := utils.GenerateRandomPassword()

	user := KeycloakUser{
		Username: email,
		Email:    email,
        Enabled: true,
        Credentials: []Credentials{
			{
				Type: "password",
				Value: randomPassword,
				Temporary: false,
			},
		},
	}
	jsonData, err := json.Marshal(user)

	if err!= nil {
		fmt.Println("ERROR: failed to marshal user data", err)
        return "", false
    }
	// Make POST api call to keycloak to create new user in master realm
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))

	if err!= nil {
        fmt.Println("ERROR: failed to create request", err)
        return "", false
    }

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + adminToken)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: failed to send request", err)
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("ERROR: failed to create user", resp)
		return "", false
	}

	fmt.Println("SUCCESS: User created successfully in keycloak", resp)

	return randomPassword, true
}

func GetKeycloakUser(adminToken, email string) (string, error) {
	apiEndpoint := "/admin/realms/master/users"
	apiURL := KeyCloakBaseUrl + apiEndpoint


	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("ERROR: failed to create request", err)
		return "", err
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR: failed to make request", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR: failed to get user", resp)
		return "", errors.New("failed to get user")
	}

	var users []struct {
        ID string `json:"id"`
    }
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		fmt.Println("ERROR: failed to decode response", err)
		return "", err
	}
	
    if len(users) == 0 {
        return "", fmt.Errorf("user not found")
    }
	return users[0].ID, nil

}

func GetKeycloakAuthToken() {

}