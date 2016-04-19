package main

import (
	"github.com/jlindsey/gobot"
	"regexp"
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
	giveReg := regexp.MustCompile(`^give (.*?){3,4}$`)
	bot.RegisterCommand(&GiveCommand{giveReg})

	bot.Start()
}
