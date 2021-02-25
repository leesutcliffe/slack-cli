package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	"net/http"
)

// Presence holds user presence value {
type Presence struct {
	Presence string `json:"presence"`
}

//TODO: duplication in cfg - refasctor
// root type for Slack user profile
type SlackProfileRoot struct {
	Profile SlackProfile `json:"profile"`
}

//TODO: duplication in cfg - refasctor
// key value pairs of user profile settings
type SlackProfile struct {
	Message    string `json:"status_text"`
	Emoji      string `json:"status_emoji"`
	Expiration int    `json:"status_expiration"`
}

// slack workspace as defined by the configuration file
type ConfigWorkspace struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}



// Do method on httpClient
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Basic api client 
type Client struct {
	token string
	httpClient httpClient
	workspace ConfigWorkspace
}

// base URL for Slack API
const baseUrl string = "https://slack.com/api/"

// returns new slack client
func New(token string) *Client {
	client := &Client{
		token: token,
		httpClient: &http.Client{},
	}

	return client
}

// SetPresence method sets the users presence to away or auto 
func (api *Client) SetPresence(value string) (string, error) {
	method := "POST"
	endpoint := "users.setPresence"
	url := baseUrl + endpoint
	data := Presence{value}
	
	req, err := userRequest(url, method, data)
	res := doPost(req, api.httpClient, api.token)
	if res.StatusCode != 200 {
		return "", err
	}
	return "", nil
}

// builds the HTTP request
func userRequest(url, method string, data interface{}) (*http.Request, error) {
	
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SetStatus method configures custom user status
func (api *Client) SetStatus(status SlackProfile) (string, error) {
	method := "POST"
	endpoint := "users.profile.set"
	url := baseUrl + endpoint
	data := SlackProfileRoot{status}

	req, err := userRequest(url, method, data)
	res := doPost(req, api.httpClient, api.token)
	if res.StatusCode != 200 {
		return "", err
	}

	return "", nil
}

//TODO: create type for req so to not rely on http.Request - better for testing
// initialtes POST request to Slack API
func doPost(req *http.Request, client httpClient, token string) *http.Response {

	auth := fmt.Sprintf("Bearer %s", token)

	req.Header.Add("Content-type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	return res
}





