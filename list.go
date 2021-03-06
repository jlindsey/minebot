package main

import (
	"github.com/jlindsey/gobot"
)

type ListCommand struct{}

func (ListCommand) String() string {
	return "ListCommand{}"
}

func (ListCommand) Help() string {
	return "*list*:  List players currently online"
}

func (ListCommand) Matches(m string) bool {
	return m == "list"
}

func (ListCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	output, err := TmuxSendKeysAndCapture(tmuxServerName, "list")

	if err != nil {
		return err
	}

	out <- gobot.NewSlackMessage(channel, stripMinecraftLogger(output))

	return nil
}
