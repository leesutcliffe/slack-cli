package main

import (
	"slack-cli/cmd"
)

type Presence struct {
	Value string `json:"presence"`
}

func main() {
	cmd.Execute()
}
