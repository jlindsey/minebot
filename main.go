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
	bot.RegisterCommand(&GiveCommand{regexp.MustCompile(`^give (.*\s?){2,}`)})

	bot.Start()
}
