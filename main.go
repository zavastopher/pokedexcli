package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(text)
	return cleanedInput
}

func commandExit() error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to close gracefully")
}

func commandHelp(commands map[string]cliCommand) error {
	helpText := "\nWelcome to the Pokedex!\nUsage:\n"
	for _, val := range commands {
		helpText += "\n" + val.name + ": " + val.description
	}
	_, err := fmt.Println("\nWelcome to the Pokedex!\n Usage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")

	if err != nil {
		return fmt.Errorf("There was an error printing help")
	}
	return nil
}
func main() {
	var commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help command",
			callback:    commandHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		commandText := cleanInput(text)[0]

		command, ok := commands[commandText]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Print(err)
			}
		}

	}
}
