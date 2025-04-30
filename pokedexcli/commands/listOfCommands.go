package commands

var ListOfCommands map[string]CliCommand

// ListOfCommands is a map of all commands available in the Pokedex CLI

func Init() {
	ListOfCommands = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the 20 next area location",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous area location",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the pokemon that you can encounter in the location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect details of a catched pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Diplays all catched pokemons",
			Callback:    CommandPokedex,
		},
	}
}
