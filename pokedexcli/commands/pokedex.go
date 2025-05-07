package commands

import "fmt"

func CommandPokedex(config *Config, data map[string]UserData, input ...string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("you didn't catch any pokemon yet")
	}

	msg := "Your Pokedex :\n"
	for key, val := range data {
		if key != val.Pokemon.Name {
			msg += "- **" + key + "** is a **" + val.Pokemon.Name + "** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
		} else {
			msg += "- **" + val.Pokemon.Name + "** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
		}
	}
	return msg, nil
}
