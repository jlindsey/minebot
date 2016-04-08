package minebot

import (
	"github.com/Jeffail/gabs"
	. "github.com/jlindsey/minebot/commands"
	"regexp"
)

type Command interface {
	SetTrigger(*regexp.Regexp)
	Help() string
	Matches(string) bool
	Run() (string, error)
}

type commandInvocation struct {
	triggeringMessage *gabs.Container
	command           Command
}

var (
	commands            []Command = make([]Command, 0)
	commandsInitialized bool      = false
)

func addCommand(s string, cmd Command) {
	trigger, err := regexp.Compile(s)
	if err != nil {
		log.Fatalf("Unable to compile regexp from string %s", s)
	}

	cmd.SetTrigger(trigger)
	commands = append(commands, cmd)
	log.Debugf(`Added command mapping for "%s": %s`, s, cmd)
}

func initCommands() {
	if commandsInitialized == true {
		return
	}

	addCommand("^ping$", &PingCommand{})

	commandsInitialized = true
}
