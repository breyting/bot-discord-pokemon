package commands

var ListOfCommands map[string]CliCommand

// ListOfCommands is a map of all commands available in the Pokedex CLI

func Init() {
	ListOfCommands = map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: ": Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: ": Displays the 20 next area location",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: ": Displays the 20 previous area location",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "[location] : Displays the pokemons that you can encounter in the location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "[pokemon] : Try to catch a pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "[pokemon] : Inspect details of a catched pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: ": Diplays all catched pokemons",
			Callback:    CommandPokedex,
		},
	}
}
