package commands

import (
	"fmt"
	"path"
	"strconv"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

func CommandMap(config *Config, data *[]UserData, input ...string) (string, error) {
	msg := "Here are the 20 next locations:\n"

	for i := 0; i < 20; i++ {
		id := path.Base(config.Next)

		location, err := GetLocation(config, id)
		if err != nil {
			return "", fmt.Errorf("error fetching location: %s", err)
		}

		msg += location.Name + "\n"

		config.Previous = config.Next
		id = path.Base(config.Previous)
		nextId, _ := strconv.Atoi(id)
		nextId += 1
		config.Next = pokeapi.BaseURL + "/location-area/" + strconv.Itoa(nextId)
	}
	return msg, nil
}
