package main

import (
	//"fmt"
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
	token := slack.GetToken("", config)

	// new slack client
	api := slack.New(token)
	_, err := api.SetPresence("auto")

	status := slack.GetStatusProfileFromConfig("", config)
	_, err = api.SetStatus(status)

	if err != nil {
		return
	}
	//fmt.Println(res)
}
