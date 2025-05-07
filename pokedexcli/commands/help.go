package commands

import (
	"fmt"
	"sort"
)

func CommandHelp(input ...string) (string, error) {
	msg := "Here is the list of the commands :\n\n"

	keys := make([]string, 0, len(ListOfCommands))

	for k := range ListOfCommands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, command := range keys {
		msg += fmt.Sprintf("- %s\n", ListOfCommands[command].Description)
	}
	return msg, nil
}
