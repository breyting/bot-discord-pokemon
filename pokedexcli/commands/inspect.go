package commands

import (
	"fmt"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

func CommandInspect(config *Config, data map[string]UserData, input ...string) (string, error) {
	pokemon := input[0]
	val, ok := data[pokemon]
	if ok {
		return printInfo(val.Pokemon), nil
	} else {
		return "", fmt.Errorf("you have not caught that pokemon")
	}
}

func printInfo(pokemonInfo pokeapi.Pokemon) string {
	msg := fmt.Sprintf("Name: %s\n", pokemonInfo.Name)
	msg += fmt.Sprintf("Height: %d\n", pokemonInfo.Height)
	msg += fmt.Sprintf("Weight: %d\n", pokemonInfo.Weight)
	msg += "Stats:\n"
	for _, val := range pokemonInfo.Stats {
		msg += fmt.Sprintf("-%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	msg += "Types:\n"
	for _, val := range pokemonInfo.Types {
		msg += fmt.Sprintf("- %s\n", val.Type.Name)
	}
	return msg
}
