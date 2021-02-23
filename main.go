package main

import (
	"fmt"
	"slack-cli/slack"
	"os/user"
)

// type SlackWorkspace struct {
// 	Name  string `yaml:"name"`
// 	Token string `yaml:"token"`
// }


// func parseConfig(configFile string) SlackConfig {

// 	data, err := ioutil.ReadFile(configFile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var config SlackConfig
// 	err = yaml.Unmarshal(data, &config)
// 	if err != nil {
// 		fmt.Printf("Error parsing YAML file: %s\n", err)
// 	}

// 	return config
// }

// func getWorkspace(workspaceName string, config SlackConfig) SlackWorkspace {
// 	var workspace SlackWorkspace

// 	for index, _ := range config.Workspace {
// 		if config.Workspace[index].Name == workspaceName {
// 			workspace = config.Workspace[index]
// 			break
// 		}
// 	}
// 	return workspace
// }

type Presence struct {
	Value string `json:"presence"`
}

func main() {

	usr, _ := user.Current()
	configFile := usr.HomeDir + "/.slack/config"

	config := slack.ParseConfig(configFile)
	// TODO: workspace can be set via a cli flag or default
	workspace := slack.GetWorkspace("cloud", config)

	p := slack.Presence{Value: "auto"}
	res, err := p.Set(workspace)
	if err != nil {
		return
	}
	fmt.Println(res)
}

// func sendRequest(method, url string, payload []byte, workspace SlackWorkspace) string {
// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	token := "Bearer "
// 	token += workspace.Token

// 	req.Header.Add("Content-type", "application/json; charset=utf-8")
// 	req.Header.Add("Authorization", token)

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)

// 	return string(body)
// }
