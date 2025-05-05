package main

import (
	"bufio"
	"fmt"
	"github.com/zavastopher/pokedexcli/internal/pokeapi.Config"
	"github.com/zavastopher/pokedexcli/internal/pokeapi.LocationResponse"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var commands map[string]cliCommand

var conf Config

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(text)
	return cleanedInput
}

func commandExit(conf *config) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to close gracefully")
}

func commandHelp(conf *config) error {
	helpText := "\nWelcome to the Pokedex!\nUsage:\n"
	for _, val := range commands {
		helpText += "\n" + val.name + ": " + val.description
	}
	_, err := fmt.Println(helpText)
	fmt.Println("")
	if err != nil {
		return fmt.Errorf("There was an error printing help")
	}
	return nil
}

func commandMap(conf *config) error {
	if conf.next == "" {
		fmt.Println("At the last page")
		return nil
	}
	var locations LocationResponse
	next, prev, err := locationsRequest(conf, &locations)
	if err != nil {
		return fmt.Errorf("Unable to get locations %v", err)
	}
	(*conf).next = next
	(*conf).previous = prev
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(conf *config) error {
	if conf.previous == "" {
		fmt.Println("Already on the first page")
		return nil
	}
	var locations LocationResponse
	next, prev, err := locationsRequestBack(conf, &locations)
	if err != nil {
		return fmt.Errorf("Unable to get locations %v", err)
	}

	(*conf).next = next
	(*conf).previous = prev
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func main() {
	commands = map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Display the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations in the Pokemon World",
			callback:    commandMapb,
		},
	}

	conf = config{
		next:     POKEAPI_ROOT_URL + "location/",
		previous: "",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		commandText := cleanInput(text)[0]
		command, ok := commands[commandText]
		if ok {
			err := command.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
