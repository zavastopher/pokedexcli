package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"internal/pokeapi"
	. "internal/pokeapi"
	"internal/pokecache"
	. "internal/pokecache"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
	Cache    Cache
}

var commands map[string]cliCommand

var conf Config

func cleanInput(text string) []string {
	cleanedInput := strings.Fields(text)
	return cleanedInput
}

func commandExit(conf *Config) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Unable to close gracefully")
}

func commandHelp(conf *Config) error {
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

func commandMap(conf *Config) error {
	if conf.Next == "" {
		fmt.Println("At the last page")
		return nil
	}

	val, ok := pokecache.Get(conf.Next, &(*conf).Cache)
	if ok {
		var locations LocationResponse
		err := json.Unmarshal(val, &locations)
		if err != nil {
			err = fmt.Errorf("Unable to unmarshal data")
		}
		for _, loc := range locations.Results {
			fmt.Println(loc.Name)
		}
		(*conf).Next = locations.Next
		if locations.Previous == nil {
			(*conf).Previous = ""
		} else {
			(*conf).Previous = *locations.Previous
		}
		return nil
	}

	var locations LocationResponse
	next, prev, results, err := pokeapi.LocationsRequest(conf.Next, &locations)
	if err != nil {
		return fmt.Errorf("Unable to get locations %v", err)
	}
	pokecache.Add(conf.Next, results, &(*conf).Cache)
	(*conf).Next = next
	(*conf).Previous = prev
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(conf *Config) error {
	if conf.Previous == "" {
		fmt.Println("Already on the first page")
		return nil
	}

	val, ok := pokecache.Get(conf.Previous, &(*conf).Cache)
	if ok {
		var locations LocationResponse
		err := json.Unmarshal(val, &locations)
		if err != nil {
			err = fmt.Errorf("Unable to unmarshal data")
		}
		for _, loc := range locations.Results {
			fmt.Println(loc.Name)
		}
		(*conf).Next = locations.Next
		if locations.Previous == nil {
			(*conf).Previous = ""
		} else {
			(*conf).Previous = *locations.Previous
		}
		return nil

	}

	var locations LocationResponse
	next, prev, results, err := pokeapi.LocationsRequest(conf.Previous, &locations)
	if err != nil {
		return fmt.Errorf("Unable to get locations %v", err)
	}
	pokecache.Add(conf.Previous, results, &(*conf).Cache)
	(*conf).Next = next
	(*conf).Previous = prev
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

	conf = Config{
		Next:     POKEAPI_ROOT_URL + "location/",
		Previous: "",
		Cache:    *NewCache(5 * time.Second),
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
