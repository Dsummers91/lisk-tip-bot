package main

import (
	"fmt"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

func main() {
	config := graw.Config{SubredditComments: []string{"lisk", "lisktip"}}
	bot, _ := reddit.NewBotFromAgentFile("tipjar.agent", 0)
	handler := &liskTipBot{bot: bot}
	if _, wait, err := graw.Run(handler, bot, config); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}
