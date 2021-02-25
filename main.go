package main

import (
	//"fmt"
	//"os/user"
	"slack-cli/slack"
	"slack-cli/cfg"
)

type Presence struct {
	Value string `json:"presence"`
}

func main() {



	//config := cfg.Parse()
	//config := cfg.New()
	config, _ := cfg.New()
	
	// TODO: workspace can be set via a cli flag or default
	token := config.GetToken("")

	// new slack client
	api := slack.New(token)
	_, err := api.SetPresence("auto")

	status := config.GetStatusProfileFromConfig("")
	_, err = api.SetStatus(status)

	if err != nil {
		return
	}
	//fmt.Println(res)
}
