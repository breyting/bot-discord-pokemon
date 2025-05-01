package commands

import (
	"errors"
	"fmt"
	"path"
	"strconv"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

func CommandMapb(config *Config, data *[]UserData, input ...string) (string, error) {
	msg := "Here are the 20 previous locations:\n"

	id := path.Base(config.Next)
	nextId, _ := strconv.Atoi(id)
	if nextId < 40 {
		return "", errors.New("can not show previous locations if not at least 40 locations are shown")
	}

	nextId -= 40
	config.Next = pokeapi.BaseURL + "/location-area/" + strconv.Itoa(nextId)
	config.Previous = pokeapi.BaseURL + "/location-area/" + strconv.Itoa(nextId-1)

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
