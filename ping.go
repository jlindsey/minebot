package main

import (
	"github.com/jlindsey/gobot"
)

// PingCommand pongs pings
type PingCommand struct{}

func (PingCommand) String() string {
	return "PingCommand{}"
}

// Run implements the Command interface
func (PingCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	out <- gobot.NewSlackMessage(channel, `Pong!`)
	return nil
}

// Matches implements the Command interface
func (PingCommand) Matches(msg string) bool {
	return msg == "ping"
}

// Help implements the Command interface
func (PingCommand) Help() string {
	return "*ping*:  Simple test command to see if I'm working."
}
