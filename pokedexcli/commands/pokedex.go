package commands

import "fmt"

func CommandPokedex(config *Config, data *[]UserData, input ...string) (string, error) {
	if len(*data) == 0 {
		return "", fmt.Errorf("you didn't catch any pokemon yet")
	}

	msg := "Your Pokedex :\n"
	for _, val := range *data {
		msg += "- **" + val.Pokedex.Name + "** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
	}
	return msg, nil
}
