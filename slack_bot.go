package main 

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxb-220695489431-VvyirohKJaga0RQ8kCMOmqGU")
	
	rtm := slack.Client.NewRTM()
	
	go rtm.ManageConnection()

	params := slack.PostMessageParameters{}
	
	channelID, timestamp, err := api.PostMessage("#general", "Some text", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
