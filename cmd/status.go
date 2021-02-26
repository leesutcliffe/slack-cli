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

func doStuff() {
	fmt.Printf("profile: %v", Profile)
	config, err := cfg.New()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	token := config.GetToken("")
	api := slack.New(token)

	status := config.GetStatusProfileFromConfig(Profile)
	res, err := api.SetStatus(status)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(res)
}
