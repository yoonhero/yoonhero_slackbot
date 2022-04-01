package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".env")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"hi",
		slack.MsgOptionText("Hello World", false),
		slack.MsgOptionAsUser(false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent successfully to channel %s at %s\n", channelID, timestamp)
}
