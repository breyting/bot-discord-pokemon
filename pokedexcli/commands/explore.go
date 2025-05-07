package commands

import (
	"errors"
	"fmt"
)

func CommandExplore(config *Config, input ...string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("can not explore without a location")
	}
	location := input[0]
	msg := fmt.Sprintf("Exploring %s...\n", location)

	locationInfo, err := GetLocation(config, location)
	if err != nil {
		return "", fmt.Errorf("error fetching location: %s", err)
	}

	msg += "Found Pokemon:\n"
	for _, val := range locationInfo.PokemonEncounters {
		msg += fmt.Sprintf("- %s\n", val.Pokemon.Name)
	}
	return msg, nil
}
