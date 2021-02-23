package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

const baseUrl string = "https://slack.com/api/"

// Presence.Set() method
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

func SetPresence(w ConfigWorkspace, value string) (string, error) {
	method := "POST"
	endpoint := "users.setPresence"
	url := baseUrl + endpoint
	presence := Presence{value}
	fmt.Printf("\n%v\n", presence)
	payload, err := json.Marshal(presence)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	return doRequest(req, w), nil

}

func SetStatus(w ConfigWorkspace, status SlackStatus) (string, error) {
	method := "POST"
	endpoint := "users.profile.set"
	url := baseUrl + endpoint
	profile := SlackProfile{status}
	fmt.Printf("\n%v\n", profile)
	payload, err := json.Marshal(profile)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	return doRequest(req, w), nil
}

func doRequest(req *http.Request, workspace ConfigWorkspace) string {
	client := &http.Client{}

	token := "Bearer "
	token += workspace.Token

	req.Header.Add("Content-type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body)
}

func GetWorkspace(workspaceName string, config ConfigRoot) ConfigWorkspace {
	if workspaceName == "" {
		workspaceName = config.Default
	}

	var workspace ConfigWorkspace

	for index, _ := range config.Workspace {
		if config.Workspace[index].Name == workspaceName {
			workspace = config.Workspace[index]
			break
		}
	}
	return workspace
}

// func GetStatus(profileName string, config ConfigRoot) ConfigStatus {
// 	var status ConfigStatus

// 	for index, _ := range config.Status {
// 		if config.Status[index].Name == profileName {
// 			status = config.Status[index]
// 			status.Expiration = 0
// 			break
// 		}
// 	}
// 	return status
// }

func GetStatus(profileName string, config ConfigRoot) SlackStatus {
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
