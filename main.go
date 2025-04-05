package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(text)
	return cleanedInput
}

func commandExit() error {
	println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to close gracefully")
}

func main() {
	type cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		command := cleanInput(text)[0]

	}
}
