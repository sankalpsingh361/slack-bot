package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main()  {
	// command line argument
	args := os.Args[1: ]
	fmt.Println(args)

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// passing the command the line arguments
	preText := "*Hello Your Jenkins build has finished*"
	jenkinsURL := "*Build URL :* " + args[0]
	var buildResult = "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*"{
		buildResult = buildResult + "  :white_check_marks: "
	} else {
		buildResult = buildResult + " :x: "
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName  + " #"  + buildNumber + "   -   "  + buildResult + "\n" + jenkinsURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)
	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		jenkinsBuildDetailsSection,
		)

	_,_,_, err := api.SendMessage(
			"C02BE77GPD3",
			msg,

		)

	if err != nil {
		fmt.Printf("%s\n",err)
		return
	}
}