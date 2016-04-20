package main

import (
	"github.com/jlindsey/gobot"
)

const (
	tmuxServerName = "minecraft"
)

func main() {
	bot := gobot.NewBot()
	gobot.StartCLI()

	bot.RegisterCommand(&PingCommand{})
	bot.RegisterCommand(&ListCommand{})
	bot.RegisterCommand(&RestartCommand{})
	bot.RegisterCommand(&GiveCommand{})
	bot.RegisterCommand(&RollCommand{})

	bot.Start()
}
