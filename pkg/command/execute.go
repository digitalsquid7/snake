package command

func Execute(commands []Command) {
	for _, command := range commands {
		command.Execute()
	}
}
