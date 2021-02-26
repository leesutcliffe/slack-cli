package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"slack-cli/cfg"
	"slack-cli/slack"
)

var cmdStatus = &cobra.Command{
	Use:   "status --profile=: Name of status profile",
	Short: "Set or get Slack profile",
	Long: `Set Slack profile using presets in the config file.
slack status 
  status:
    - name: lunch
    message: "Out to lunch"
    emoji: ":bagel:"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		doStuff()
	},
}

// temp function
//TODO: break up into separate methods so token can be resused 
func doStuff() {
	
	config, err := cfg.New()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	token := config.GetToken("")
	api := slack.New(token)

	if Away != false {
		res, err := api.SetPresence("away")
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		fmt.Println(res)
	} else {
		res, err := api.SetPresence("auto")
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		fmt.Println(res)
	}

	status := config.GetStatusProfileFromConfig(Profile)
	_, err = api.SetStatus(status)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

}

