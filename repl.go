package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	commands "github.com/breyting/pokedex-discord/pokedexcli/commands"
	pokecache "github.com/breyting/pokedex-discord/pokedexcli/pokecache"
)

var listOfCommands map[string]commands.CliCommand
var conf commands.Config
var cache = pokecache.NewCache((5 * time.Second))

func startRepl(conf *commands.Config) {
	conf.Next = "https://pokeapi.co/api/v2/location-area/1"

	listOfCommands = map[string]commands.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the 20 next area location",
			Callback:    commands.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous area location",
			Callback:    commands.CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the pokemon that you can encounter in the location",
			Callback:    commands.CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon",
			Callback:    commands.CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect details of a catched pokemon",
			Callback:    commands.CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Diplays all catched pokemons",
			Callback:    commands.CommandPokedex,
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
			err := command.Callback(conf, input[1:])
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
