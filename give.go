package main

import (
	"fmt"
	"github.com/jlindsey/gobot"
	"regexp"
)

type GiveCommand struct {
	matcher *regexp.Regexp
}

func (g GiveCommand) String() string {
	return fmt.Sprintf("GiveCommand{ matcher: %s }", g.matcher.String())
}

func (GiveCommand) Help() string {
	return `*give*: Give an item to the specified player.
	This takes the same arguments as the /give command in Minecraft. For example:
	
	Gives notch 1 Carrot (by block ID or name)
	@minebot: give notch 392 1
	@minebot: give notch minecraft:carrot 1

	Gives notch 10 Yellow wool (using block subids)
	@minebot: give notch 35 10 4
	@minebot: give notch minecraft:wool 10 4`
}

func (g GiveCommand) Matches(str string) bool {
	return g.matcher.MatchString(str)
}

func (g *GiveCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	return nil
}
