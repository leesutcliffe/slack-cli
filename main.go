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
	//token := workspace.Token

	// new slack client
	api := slack.New(token)
	_, err := api.SetPresence("auto")

	
	// status := slack.GetStatus("", config)
	// res, err = slack.SetStatus(workspace, status)


	//_, err := slack.SetPresence(workspace, "auto")
	
	if err != nil {
		return
	}
	//fmt.Println(res)
}
