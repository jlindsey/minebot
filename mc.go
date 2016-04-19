package main

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	mcLoggerReg = regexp.MustCompile(`^\[\d+:\d+:\d+\] \[.*?\]: (.*)$`)
)

func stripMinecraftLogger(str string) string {
	var buf bytes.Buffer

	for _, line := range strings.Split(str, "\n") {
		buf.WriteString(mcLoggerReg.ReplaceAllString(line, "$1"))
	}

	return buf.String()
}
