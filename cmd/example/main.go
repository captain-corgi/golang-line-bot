package main

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	fmt.Printf("Hello Line Friends\n")

	req := http.Request{}

	// Initial Client
	client := &http.Client{}
	bot, err := linebot.New("<channel secret>", "<channel accsss token>", linebot.WithHTTPClient(client))

	// Parse request
	events, err := bot.ParseRequest(&req)
	if err != nil {
		// Do something when something bad happened.
	}

	// Loop for events
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			// Do Something...
		}
	}

}
