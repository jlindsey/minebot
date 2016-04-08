package minebot

import "regexp"

type PingCommand struct{}

func (PingCommand) String() string {
	return "PingCommand{}"
}

func (*PingCommand) SetTrigger(r *regexp.Regexp) {}

func (PingCommand) Run() (string, error) {
	return "Pong", nil
}

func (*PingCommand) Matches(msg string) bool {
	return msg == "ping"
}

func (PingCommand) Help() string {
	return "*ping*:  Simple test command to see if I'm working."
}
