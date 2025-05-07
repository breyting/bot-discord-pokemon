package commands

import "fmt"

func CommandPokedex(config *Config, data map[string]UserData, input ...string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("you didn't catch any pokemon yet")
	}

	if input[0] == "shiny" {
		msg := "Your shiny Pokedex :\n"
		for key, val := range data {
			if val.IsShiny {
				if key != val.Pokemon.Name {
					msg += "- **" + key + "** is a **__shiny " + val.Pokemon.Name + "__** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
				} else {
					msg += "- **" + val.Pokemon.Name + "** is a **__shiny pokemon__** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
				}
			}
		}
		return msg, nil
	}

	msg := "Your Pokedex :\n"
	for key, val := range data {
		if key != val.Pokemon.Name {
			if val.IsShiny {
				msg += "- **" + key + "** is a **__shiny " + val.Pokemon.Name + "__** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
			} else {
				msg += "- **" + key + "** is a **" + val.Pokemon.Name + "** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
			}
		} else {
			if val.IsShiny {
				msg += "- **" + val.Pokemon.Name + "** is a **__shiny pokemon__** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
			} else {
				msg += "- **" + val.Pokemon.Name + "** captured the " + val.CaptureDate.Format("2006-01-02 15:04:05") + "\n"
			}
		}
	}
	return msg, nil
}
