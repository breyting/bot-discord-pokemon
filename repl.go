package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	commands "github.com/breyting/pokedex-discord/pokedexcli/commands"
	pokecache "github.com/breyting/pokedex-discord/pokedexcli/internal/pokecache"
)

var listOfCommands map[string]commands.CliCommand
var conf config
var cache = pokecache.NewCache((5 * time.Second))

func start_repl() {
	conf = config{
		next:     "https://pokeapi.co/api/v2/location-area/1",
		previous: "",
	}

	listOfCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commands.CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commands.CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the 20 next area location",
			callback:    commands.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 previous area location",
			callback:    commands.CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon that you can encounter in the location",
			callback:    commands.CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commands.CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect details of a catched pokemon",
			callback:    commands.CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Diplays all catched pokemons",
			callback:    commands.CommandPokedex,
		},
	}

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		text := scan.Text()
		input := cleanInput(text)

		if len(input) == 0 {
			continue
		}

		command, exists := listOfCommands[input[0]]
		if exists {
			err := command.callback(&conf, input[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
