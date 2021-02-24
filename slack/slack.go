package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type Presence struct {
	Presence string `json:"presence"`
}

type SlackProfile struct {
	Profile SlackStatus `json:"profile"`
}

type SlackStatus struct {
	Message    string `json:"status_text"`
	Emoji      string `json:"status_emoji"`
	Expiration int    `json:"status_expiration"`
}

type ConfigWorkspace struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

type ConfigRoot struct {
	Default   string            `yaml:"default"`
	Workspace []ConfigWorkspace `yaml:"workspaces"`
	Status    []ConfigStatus    `yaml:"status"`
}

type ConfigStatus struct {
	Name       string `yaml:"name"`
	Message    string `yaml:"message"`
	Emoji      string `yaml:"emoji"`
	Expiration int
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	token string
	httpClient httpClient
	workspace ConfigWorkspace
}

const baseUrl string = "https://slack.com/api/"

func New(token string) *Client {
	client := &Client{
		token: token,
		httpClient: &http.Client{},
	}

	return client
}

func (api *Client) SetPresence(value string) (string, error) {
	method := "POST"
	endpoint := "users.setPresence"
	url := baseUrl + endpoint
	data := Presence{value}
	
	req, err := userRequest(url, method, data)
	res := api.doPost(req)
	if res.StatusCode != 200 {
		return "", err
	}
	return "", nil
}

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

func (api *Client) SetStatus(status SlackStatus) (string, error) {
	method := "POST"
	endpoint := "users.profile.set"
	url := baseUrl + endpoint
	data := SlackProfile{status}

	req, err := userRequest(url, method, data)
	res := api.doPost(req)
	if res.StatusCode != 200 {
		return "", err
	}

	return "", nil
}

func (api *Client) doPost(req *http.Request) *http.Response {
	client := &http.Client{}

	token := "Bearer "
	token += api.token

	req.Header.Add("Content-type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)

	return res
}


func GetToken(workspaceName string, config ConfigRoot) string {
	if workspaceName == "" {
		workspaceName = config.Default
	}

	var workspace string

	for index, _ := range config.Workspace {
		if config.Workspace[index].Name == workspaceName {
			workspace = config.Workspace[index].Token
			break
		}
	}
	return workspace
}

func GetStatusProfileFromConfig(profileName string, config ConfigRoot) SlackStatus {
	var status SlackStatus

	for index, _ := range config.Status {
		if config.Status[index].Name == profileName {
			status.Emoji = config.Status[index].Emoji
			status.Message = config.Status[index].Message
			status.Expiration = 0
			break
		}
	}
	return status
}

func ParseConfig(configFile string) ConfigRoot {

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
	}
	var config ConfigRoot
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	return config
}
