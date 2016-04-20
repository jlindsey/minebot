package main

import (
	"fmt"
	"github.com/jlindsey/gobot"
	"regexp"
)

var (
	giveReg = regexp.MustCompile(`^give (.*?){3,4}$`)
)

// GiveCommand gives items to players
type GiveCommand struct{}

func (g GiveCommand) String() string {
	return fmt.Sprintf("GiveCommand{ matcher: %s }", giveReg.String())
}

// Help implements the Command interface
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

// Matches implements the Command interface
func (g GiveCommand) Matches(str string) bool {
	return giveReg.MatchString(str)
}

// Run implements the Command interface
func (g *GiveCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	output, err := TmuxSendKeysAndCapture(tmuxServerName, text)

	if err != nil {
		return err
	}

	out <- gobot.NewSlackMessage(channel, stripMinecraftLogger(output))

	return nil
}
