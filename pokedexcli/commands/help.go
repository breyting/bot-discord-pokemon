package commands

import "fmt"

func CommandHelp(config *Config, input []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage :")
	fmt.Println()
	for _, command := range listOfCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
