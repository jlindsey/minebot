out/minebot: minebot/main.go minebot.go log.go tmux.go commands.go message.go cli.go commands/ping.go
	cd minebot && go build -o ../out/minebot main.go

clean:
	rm -rf out
	rm -rf *.log

run: out/minebot
	./out/minebot

.PHONY: clean run
