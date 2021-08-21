package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main()  {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C02BE77GPD3",
		slack.MsgOptionText("hello world!", false),
		)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message send seccessfully to channel %s at %s",channelID,timestamp)

}