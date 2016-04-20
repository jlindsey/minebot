package main

import (
	"fmt"
	"github.com/jlindsey/gobot"
	"time"
)

// RestartCommand restarts the server after a 1-minute interval
type RestartCommand struct {
	inProgress bool
}

func (r RestartCommand) String() string {
	return fmt.Sprintf("RestartCommand{ inProgress: %s }", r.inProgress)
}

// Help implements the Command interface
func (RestartCommand) Help() string {
	return "*restart*: Restart the server after a 1-minute countdown"
}

// Matches implements the Command interface
func (RestartCommand) Matches(msg string) bool {
	return msg == "restart"
}

// Run implements the Command interface
func (r *RestartCommand) Run(channel string, text string, out chan *gobot.SlackMessage) error {
	if r.inProgress {
		out <- gobot.NewSlackMessage(channel, "A restart is already in progress!")
		return nil
	}

	go r.doRestart(channel, out)

	return nil
}

func (r *RestartCommand) doRestart(channel string, out chan *gobot.SlackMessage) {
	r.inProgress = true

	slackAndServer := func(msg string) {
		out <- gobot.NewSlackMessage(channel, msg)
		TmuxSendKeys(tmuxServerName, fmt.Sprintf("say %s", msg))
	}

	done := make(chan bool)

	slackAndServer("Restarting in 1 minute")
	time.AfterFunc(30*time.Second, func() { slackAndServer("Restarting in 30 seconds") })
	time.AfterFunc(45*time.Second, func() { slackAndServer("Restarting in 15 seconds") })
	time.AfterFunc(55*time.Second, func() { slackAndServer("Restarting in 5 seconds") })
	time.AfterFunc(56*time.Second, func() { slackAndServer("Restarting in 4 seconds") })
	time.AfterFunc(57*time.Second, func() { slackAndServer("Restarting in 3 seconds") })
	time.AfterFunc(58*time.Second, func() { slackAndServer("Restarting in 2 seconds") })
	time.AfterFunc(59*time.Second, func() { slackAndServer("Restarting in 1 second") })
	time.AfterFunc(60*time.Second, func() {
		slackAndServer("Restarting NOW")
		TmuxSendKeys(tmuxServerName, "stop")
		done <- true
	})

	<-done
	r.inProgress = false
}
