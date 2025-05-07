package commands

import (
	"time"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
}

type Config struct {
	PokeapiClient pokeapi.Client
	Next          string
	Previous      string
	CaughtPokemon map[string]pokeapi.Pokemon
}

// ListOfCommands is a map of all commands available in the Pokedex CLI
var ListOfCommands map[string]CliCommand

func Init() {
	ListOfCommands = map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "`help` : Displays a help message",
		},
		"map": {
			Name:        "map",
			Description: "`map` : Displays the 20 next area location",
		},
		"mapb": {
			Name:        "mapb",
			Description: "`mapb` : Displays the 20 previous area location",
		},
		"explore": {
			Name:        "explore",
			Description: "`explore [location]` : Displays the pokemons that you can encounter in the location",
		},
		"catch": {
			Name:        "catch",
			Description: "`catch [pokemon]` : Try to catch a pokemon",
		},
		"inspect": {
			Name:        "inspect",
			Description: "`inspect [pokemon]` : Inspect details of a catched pokemon",
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "`pokedex` : Diplays all catched pokemons",
		},
	}
}

type UserData struct {
	Pokemon     pokeapi.Pokemon `json:"pokedex"`
	CaptureDate time.Time       `json:"capture_date"`
	IsShiny     bool            `json:"is_shiny"`
}
