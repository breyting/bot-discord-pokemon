package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

func CommandExplore(config *Config, data *[]UserData, input ...string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("Can not explore without a location")
	}
	location := input[0]
	msg := fmt.Sprintf("Exploring %s...\n", location)

	cache := config.PokeapiClient.Cache
	val, ok := cache.Get(location)
	if ok {
		var locationArea pokeapi.Location
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return "", fmt.Errorf("Unmarshal error : %s", err)
		}

		msg += "Found Pokemon:\n"
		for _, val := range locationArea.PokemonEncounters {
			msg += fmt.Sprintf("- %s\n", val.Pokemon.Name)
		}
		return msg, nil
	} else {
		api_location := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
		res, err := http.Get(api_location)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return "", errors.New("this location doesn't exist")
		}

		var locationArea pokeapi.Location
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(&locationArea); err != nil {
			msg := fmt.Sprintf("Decoder error : %s", err)
			return "", errors.New(msg)
		}

		msg += "Found Pokemon:\n"
		for _, val := range locationArea.PokemonEncounters {
			msg += fmt.Sprintf("- %s\n", val.Pokemon.Name)
		}
		data, err := json.Marshal(locationArea)
		if err != nil {
			return "", fmt.Errorf("error with Marshal: %s", err)
		}
		cache.Add(location, []byte(data))
		return msg, nil
	}
}
