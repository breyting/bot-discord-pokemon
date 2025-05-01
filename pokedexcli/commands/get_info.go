package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

func GetPokemon(config *Config, input string) (pokeapi.Pokemon, error) {
	var pokemonInfo pokeapi.Pokemon

	cache := config.PokeapiClient.Cache
	val, ok := cache.Get(input)
	if ok {
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			msg := fmt.Sprintf("Unmarshal error : %s", err)
			return pokeapi.Pokemon{}, errors.New(msg)
		}
	} else {
		api_pokemon := pokeapi.BaseURL + "/pokemon/" + input

		res, err := http.Get(api_pokemon)
		if err != nil {
			return pokeapi.Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return pokeapi.Pokemon{}, errors.New("this pokemon doesn't exist")
		}

		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&pokemonInfo); err != nil {
			return pokeapi.Pokemon{}, fmt.Errorf("decoder error : %s", err)
		}

		data, err := json.Marshal(pokemonInfo)
		if err != nil {
			return pokeapi.Pokemon{}, fmt.Errorf("error with Marshal: %s", err)
		}
		cache.Add(input, []byte(data))
	}

	return pokemonInfo, nil
}

func GetLocation(config *Config, input string) (pokeapi.Location, error) {
	var locationArea pokeapi.Location

	cache := config.PokeapiClient.Cache
	val, ok := cache.Get(input)
	if ok {
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			msg := fmt.Sprintf("Unmarshal error : %s", err)
			return pokeapi.Location{}, errors.New(msg)
		}
	} else {
		api_location := pokeapi.BaseURL + "/location-area/" + input

		res, err := http.Get(api_location)
		if err != nil {
			return pokeapi.Location{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return pokeapi.Location{}, errors.New("this location doesn't exist")
		}

		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&locationArea); err != nil {
			msg := fmt.Sprintf("Decoder error : %s", err)
			return pokeapi.Location{}, errors.New(msg)
		}

		data, err := json.Marshal(locationArea)
		if err != nil {
			return pokeapi.Location{}, fmt.Errorf("error with Marshal: %s", err)
		}
		cache.Add(input, []byte(data))
	}

	return locationArea, nil
}
