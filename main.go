package main

import (
	"fmt"
	"os/user"
	"slack-cli/slack"
)

type Presence struct {
	Value string `json:"presence"`
}

func main() {

	usr, _ := user.Current()
	configFile := usr.HomeDir + "/.slack/config"

	config := slack.ParseConfig(configFile)
	// TODO: workspace can be set via a cli flag or default
	workspace := slack.GetWorkspace("", config)
	status := slack.GetStatus("", config)


	res, err := slack.SetPresence(workspace, "auto")
	
	res, err = slack.SetStatus(workspace, status)
	if err != nil {
		return
	}
	fmt.Println(res)
}
