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

func CommandCatch(config *Config, input []string) error {
	if len(input) == 0 {
		return errors.New("Can not catch without a pokemon name")
	}
	pokemonInput := input[0]
	cache := config.PokeapiClient.Cache
	val, ok := cache.Get(pokemonInput)
	if ok {
		var pokemonInfo pokeapi.Pokemon
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			msg := fmt.Sprintf("Unmarshal error : %s", err)
			return errors.New(msg)
		}
		tryingCatch(pokemonInfo)
	} else {
		api_pokemon := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonInput)
		res, err := http.Get(api_pokemon)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			msg := fmt.Sprintf("This pokemon doesn't exist!")
			return errors.New(msg)
		}
		var pokemonInfo pokeapi.Pokemon
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&pokemonInfo); err != nil {
			msg := fmt.Sprintf("Decoder error : %s", err)
			return errors.New(msg)
		}

		data, err := json.Marshal(pokemonInfo)
		if err != nil {
			return fmt.Errorf("error with Marshal: %s", err)
		}
		cache.Add(pokemonInput, []byte(data))

		tryingCatch(pokemonInfo)
	}
	return nil
}

func tryingCatch(pokemonInfo pokeapi.Pokemon) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	baseExperience := pokemonInfo.BaseExperience
	chance := rand.Intn(baseExperience)
	if chance < 50 {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
		ownPokedex[pokemonInfo.Name] = pokemonInfo
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}
}
