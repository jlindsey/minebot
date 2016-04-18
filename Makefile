SOURCES := $(shell find . -name "*.go")

minebot: ${SOURCES}
	go build

run: minebot
	./minebot

clean:
	go clean

.PHONY: run clean
