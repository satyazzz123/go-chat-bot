package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
)

func welcome(event *irc.Event) {
	fmt.Printf("Welcome, %s!\n", event.Nick)
}

func main() {
	ircCon := irc.IRC("YourBotName", "YourBotName")
	ircCon.Connect("irc.freenode.net:6667")

	ircCon.AddCallback("001", func(e *irc.Event) {
		ircCon.Join("#chatbot_test_channel")
	})

	ircCon.AddCallback("PRIVMSG", func(e *irc.Event) {
		if e.Arguments[0] == "#chatbot_test_channel" {
			welcome(e)
		}
	})

	ircCon.Loop()
}
