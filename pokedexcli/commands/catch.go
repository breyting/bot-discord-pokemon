package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

var ownPokedex = map[string]pokeapi.Pokemon{}

func CommandCatch(config *Config, data *[]UserData, input ...string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("can not catch without a pokemon name")
	}

	pokemonInput := input[0]

	pokemonInfo, err := GetPokemon(config, pokemonInput)
	if err != nil {
		return "", err
	}

	return tryingCatch(pokemonInfo, data)
}

func tryingCatch(pokemonInfo pokeapi.Pokemon, data *[]UserData) (string, error) {
	baseExperience := pokemonInfo.BaseExperience
	chance := rand.Intn(baseExperience)
	msg := fmt.Sprintf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	if chance < 50 {
		ownPokedex[pokemonInfo.Name] = pokemonInfo
		new_pokemon := UserData{
			pokemonInfo,
			time.Now(),
		}
		*data = append(*data, new_pokemon)
		sprite := pokemonInfo.Sprites.FrontDefault
		msg += fmt.Sprintf("%s was caught!\n", pokemonInfo.Name)
		msg += fmt.Sprintf("%v\n", sprite)
		msg += "You may now inspect it with the inspect command.\n"
		return msg, nil
	} else {
		sprite := pokemonInfo.Sprites.BackDefault
		msg += fmt.Sprintf("%s escaped!\n", pokemonInfo.Name)
		msg += fmt.Sprintf("%v\n", sprite)
		return msg, nil
	}
}
