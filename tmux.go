package minebot

import (
	"bytes"
	"fmt"
	"os/exec"
)

const tmuxServerName = "minecraft"

var commandSem = make(chan int, 1)

func tmuxSendKeysAndCapture(keys string) (*bytes.Buffer, error) {
	defaultArgs := []string{
		"-L",
		tmuxServerName,
		"send-keys",
	}
	args := make([]string, 5)
	var (
		cmd *exec.Cmd
		err error
		buf bytes.Buffer
	)

	copy(args, defaultArgs)
	args = append(args, fmt.Sprintf(`"%s"`, keys), `"Enter"`)

	commandSem <- 1
	cmd = exec.Command("tmux", args...)
	err = cmd.Run()
	if err != nil {
		log.Errorf("Error running command: %s", err)
		return nil, nil
	}

	cmd = exec.Command("tmux", "-L", tmuxServerName, "capture-pane")
	err = cmd.Run()
	if err != nil {
		log.Errorf("Error running command: %s", err)
	}

	cmd = exec.Command("tmux", "-L", tmuxServerName, "show-buffer")
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		log.Errorf("Error running command: %s", err)
	}
	<-commandSem

	return &buf, nil
}
