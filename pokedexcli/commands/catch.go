package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

var ownPokedex = map[string]pokeapi.Pokemon{}

func CommandCatch(config *Config, input ...string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("can not catch without a pokemon name")
	}
	pokemonInput := input[0]
	cache := config.PokeapiClient.Cache
	val, ok := cache.Get(pokemonInput)
	if ok {
		var pokemonInfo pokeapi.Pokemon
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			msg := fmt.Sprintf("Unmarshal error : %s", err)
			return "", errors.New(msg)
		}
		return tryingCatch(pokemonInfo)
	} else {
		api_pokemon := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonInput)
		res, err := http.Get(api_pokemon)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return "", errors.New("this pokemon doesn't exist")
		}
		var pokemonInfo pokeapi.Pokemon
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&pokemonInfo); err != nil {
			msg := fmt.Sprintf("Decoder error : %s", err)
			return "", errors.New(msg)
		}

		data, err := json.Marshal(pokemonInfo)
		if err != nil {
			return "", fmt.Errorf("error with Marshal: %s", err)
		}
		cache.Add(pokemonInput, []byte(data))

		return tryingCatch(pokemonInfo)
	}
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
