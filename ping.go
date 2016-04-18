package main

import (
	"github.com/jlindsey/gobot"
)

type PingCommand struct{}

func (PingCommand) String() string {
	return "PingCommand{}"
}

func (PingCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	out <- gobot.NewSlackMessage(channel, `Pong!`)
	return nil
}

func (PingCommand) Matches(msg string) bool {
	return msg == "ping"
}

func (PingCommand) Help() string {
	return "*ping*:  Simple test command to see if I'm working."
}
