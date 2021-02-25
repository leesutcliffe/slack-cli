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
	// TODO: workspace can be set via a cli flag or default
	token := cfg.GetToken("")

	// new slack client
	api := slack.New(token)
	_, err := api.SetPresence("auto")

	status := cfg.GetStatusProfileFromConfig("")
	_, err = api.SetStatus(status)

	if err != nil {
		return
	}
	//fmt.Println(res)
}
