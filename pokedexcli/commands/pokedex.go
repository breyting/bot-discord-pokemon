package commands

import "fmt"

func CommandPokedex(config *Config, input ...string) (string, error) {
	if len(ownPokedex) == 0 {
		return "", fmt.Errorf("you didn't catch any pokemon yet")
	}

	msg := "Your Pokedex :\n"
	for _, val := range ownPokedex {
		msg += fmt.Sprintf("- %s\n", val.Name)
	}
	return msg, nil
}
