package main

import (
	"fmt"
	"slack-cli/slack"
	"os/user"
)

type Presence struct {
	Value string `json:"presence"`
}

func main() {

	usr, _ := user.Current()
	configFile := usr.HomeDir + "/.slack/config"

	config := slack.ParseConfig(configFile)
	// TODO: workspace can be set via a cli flag or default
	workspace := slack.GetWorkspace("cloud", config)

	//presence := slack.Presence{Presence: "auto"}
	res, err := slack.SetPresence(workspace, "auto")
	if err != nil {
		return
	}
	fmt.Println(res)
}