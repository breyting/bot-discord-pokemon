package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

var ownPokedex = map[string]pokeapi.Pokemon{}

func CommandCatch(config *Config, input ...string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("can not catch without a pokemon name")
	}

	pokemonInput := input[0]

	pokemonInfo, err := GetPokemon(config, pokemonInput)
	if err != nil {
		return "", err
	}

	return tryingCatch(pokemonInfo)
}

func tryingCatch(pokemonInfo pokeapi.Pokemon) (string, error) {
	baseExperience := pokemonInfo.BaseExperience
	chance := rand.Intn(baseExperience)
	if chance < 50 {
		ownPokedex[pokemonInfo.Name] = pokemonInfo
		return fmt.Sprintf("Throwing a Pokeball at %s...\n%s was caught!\nYou may now inspect it with the inspect command.", pokemonInfo.Name, pokemonInfo.Name), nil
	} else {
		return fmt.Sprintf("Throwing a Pokeball at %s...\n%s escaped!\n", pokemonInfo.Name, pokemonInfo.Name), nil
	}
}
