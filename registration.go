package main

func CommandRegistration() {
	RegisterCommand("vulcan", "Show or modify vulcan instance configurations.", []string{"subcommand"}, VulcanCommand)
	RegisterCommand("cd", "Change your current working directory.", []string{"dir"}, CdCommand)
	RegisterCommand("ls", "List all files in a directory.", []string{}, LsCommand)
	RegisterCommand("help", "Show some help about the commands.", []string{}, HelpCommand)
	RegisterCommand("clear", "Clear the console window.", []string{}, ClearCommand)
	RegisterCommand("exit", "Exit the application.", []string{}, ExitCommand)
}

func RegisterCommand(name string, description string, args []string, runnable func([]string)) {
	arguments := Args{
		Count: len(args),
		Get:   args,
	}
	commands = append(commands, Command{
		Name:        name,
		Description: description,
		Args:        arguments,
		Run:         runnable,
	})
}
