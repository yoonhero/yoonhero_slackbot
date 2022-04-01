package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/yoonhero/yoonhero_slackbot/messages"
)

func main() {
	_ = godotenv.Load(".env")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	err := messages.SendMsg("HI", api)

	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err)
	}
}
