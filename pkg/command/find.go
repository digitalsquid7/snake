package command

func Find(finders []Finder) []Command {
	commands := make([]Command, 0)
	for _, finder := range finders {
		if command := finder.Find(); command != nil {
			commands = append(commands, command)
		}
	}
	return commands
}
