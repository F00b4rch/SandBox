package slackApp

import (
	"fmt"

	"github.com/ashwanthkumar/slack-go-webhook"
)

// PayloadSlack create an app in slack API and add an Incoming Webhoosk
// more informations : https://api.slack.com/apps/
func PayloadSlack(text string) {

	// Add here your webhookurl
	webhookURL := "https://hooks.slack.com/services/XXXX-XXXXX-XXXXX-XXXX"

	payload := slack.Payload{
		Text:     text,
		Username: "User",
		Channel:  "#infra",
	}
	err := slack.Send(webhookURL, "", payload)
	if len(err) > 0 {
		fmt.Printf("error : %s\n", err)
	}
}
