package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var commandSem = make(chan int, 1)

func randHash() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	c := md5.Sum(b)
	return fmt.Sprintf("%x", c), nil
}

func getDelimiters(hash string) (string, string) {
	begin := fmt.Sprintf("### %s ###", hash)
	end := fmt.Sprintf("###/ %s ###", hash)

	return begin, end
}

func commandsForOperation(tmuxServerName string, keys string) (hash string, commands [][]string, err error) {
	hash, err = randHash()
	if err != nil {
		return
	}

	begin, end := getDelimiters(hash)

	commands = [][]string{
		[]string{"-L", tmuxServerName, "send-keys", begin, "Enter"},
		[]string{"-L", tmuxServerName, "send-keys", keys, "Enter"},
		[]string{"-L", tmuxServerName, "send-keys", end, "Enter"},
		[]string{"-L", tmuxServerName, "capture-pane"},
		[]string{"SLEEP"},
		[]string{"-L", tmuxServerName, "show-buffer"},
	}

	return
}

func parseOutput(str string, hash string) (s string, err error) {
	begin, end := getDelimiters(hash)

	startI := int64(strings.Index(str, begin))
	if startI == -1 {
		err = fmt.Errorf("Unable to find start delimiter in tmux output")
		return
	}

	startI = startI + int64(len(begin))
	endI := int64(strings.Index(str, end))
	if endI == -1 {
		err = fmt.Errorf("Unable to find end delimiter in tmux output")
		return
	}
	length := endI - startI

	b := make([]byte, length)
	r := strings.NewReader(str)
	r.Seek(startI, 0)
	i, err := r.Read(b)

	s = strings.SplitN(string(b[:i]), "\n", 4)[3]
	return
}

// TmuxSendKeys runs a command inside a detached tmux client
func TmuxSendKeys(tmuxServerName string, keys string) error {
	commandSem <- 1
	cmd := exec.Command("tmux", "-L", tmuxServerName, "send-keys", keys, "Enter")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error running command: %s", err)
	}
	<-commandSem

	return nil
}

// TmuxSendKeysAndCapture runs a command inside a detached tmux client and returns the buffer output.
func TmuxSendKeysAndCapture(tmuxServerName string, keys string) (str string, err error) {
	var (
		hash     string
		commands [][]string
		buf      bytes.Buffer
	)

	commandSem <- 1
	hash, commands, err = commandsForOperation(tmuxServerName, keys)
	if err != nil {
		return
	}

	for i, args := range commands {
		if len(args) == 1 && args[0] == "SLEEP" {
			time.Sleep(500 * time.Millisecond)
		}

		cmd := exec.Command("tmux", args...)

		if (i + 1) == len(commands) {
			cmd.Stdout = &buf
		}

		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("Error running command: %s", err)
			return
		}
	}
	<-commandSem

	str, err = parseOutput(buf.String(), hash)
	return
}
