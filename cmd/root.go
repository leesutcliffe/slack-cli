package cmd

import (
	"github.com/spf13/cobra"
)

var Profile string
var Away bool

func init() {
	cmdStatus.Flags().StringVarP(&Profile, "profile", "p", "", "name of staus profile from config")
	cmdStatus.Flags().BoolVarP(&Away, "away", "a", false, "sets status to away, default: False")
}

func Execute() {

	var cmdSet = &cobra.Command{
		Use: "set [resource to set]",
	}

	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "slack"}
	rootCmd.AddCommand(cmdSet)
	cmdSet.AddCommand(cmdStatus)
	rootCmd.Execute()
}
