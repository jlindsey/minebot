package minebot

type ListCommand struct{}

func (ListCommand) String() string {
	return "ListCommand{}"
}

func (ListCommand) Help() string {
	return "*list*:  List players currently online"
}

func (ListCommand) Matches(m string) bool {
	return m == "list"
}

func (ListCommand) Run() (out string, err error) {
	return tmuxSendKeysAndCapture("list")
}
