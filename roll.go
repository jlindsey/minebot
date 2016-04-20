package main

import (
	"fmt"
	"github.com/jlindsey/gobot"
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"
)

var (
	rollParser = regexp.MustCompile(`(?i)^roll (\d+)d(\d+)(\+|\-)?(\d+)?$`)
)

const (
	minRoll = 1
)

const (
	modeSimple  = 2
	modeWithMod = 4
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RollCommand rolls various dice
type RollCommand struct{}

func (RollCommand) String() string {
	return fmt.Sprintf("RollCommand{ parser: %s }", rollParser.String())
}

// Matches implements the Command interface
func (RollCommand) Matches(text string) bool {
	return rollParser.MatchString(text)
}

// Help implements the Command interface
func (RollCommand) Help() string {
	return `*roll*: Roll some dice. Using normal D&D nomenclature, you can roll dice with modifiers. Examples:
		@minebot: roll 1d20
		@minebot: roll 3d6+5`
}

// Run implements the Command interface
func (*RollCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	var result int
	var msgs []string

	createMessage := func(msg string) {
		msgs = append(msgs, msg)
	}

	matches := rollParser.FindStringSubmatch(text)
	matches = filterBlank(matches[1:])

	mode := len(matches)
	if mode != modeSimple && mode != modeWithMod {
		out <- gobot.NewSlackMessage(channel, "Bad roll input!")
		return fmt.Errorf("Bad roll input: %s", text)
	}

	num, die, op, mod, err := convertMatches(matches)
	if err != nil {
		out <- gobot.NewSlackMessage(channel, err.Error())
		return err
	}

	for i := 0; i < num; i++ {
		dieRoll := minRand(die)
		createMessage(fmt.Sprintf(":d20: %d", dieRoll))
		result += dieRoll
	}

	if mode == modeWithMod {
		switch op {
		case '+':
			createMessage(fmt.Sprintf(":heavy_plus_sign:%d", mod))
			result += mod
		case '-':
			result -= mod
			createMessage(fmt.Sprintf(":heavy_minus_sign:%d", mod))
		default:
			out <- gobot.NewSlackMessage(channel, "Only `+` and `-` modifiers allowed!")
			return fmt.Errorf("Only `+` and `-` modifiers allowed!")
		}
	}

	if num > 1 || mod > 0 {
		createMessage(fmt.Sprintf("> `%d`", result))
	}

	for _, msg := range msgs {
		out <- gobot.NewSlackMessage(channel, msg)
		// Sleep between pushing on the channel otherwise
		// the messages get sent out of order
		time.Sleep(5 * time.Millisecond)
	}

	return nil
}

func filterBlank(matches []string) []string {
	out := make([]string, 0, 4)

	for _, str := range matches {
		if len(str) > 0 {
			out = append(out, str)
		}
	}

	return out
}

func convertMatches(matches []string) (num int, die int, op rune, mod int, err error) {
	num, err = strconv.Atoi(matches[0])
	die, err = strconv.Atoi(matches[1])

	if num < 1 {
		err = fmt.Errorf("Number of dice must be one or more!")
		return
	}

	if die < 2 {
		err = fmt.Errorf("There's no such thing as a %d-sided die!", die)
		return
	}

	if len(matches) == 4 {
		op, _ = utf8.DecodeRuneInString(matches[2])
		mod, err = strconv.Atoi(matches[3])
	}

	return
}

func minRand(i int) int {
	return rand.Intn(i-minRoll) + minRoll
}
