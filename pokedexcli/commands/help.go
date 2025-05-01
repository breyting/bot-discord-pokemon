package commands

import (
	"fmt"
	"sort"
)

func CommandHelp(config *Config, input ...string) (string, error) {
	msg := "Here is the list of the commands :\n\n"

	keys := make([]string, 0, len(ListOfCommands))

	for k := range ListOfCommands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, command := range keys {
		msg += fmt.Sprintf("- `%s` %s\n", ListOfCommands[command].Name, ListOfCommands[command].Description)
	}
	return msg, nil
}
