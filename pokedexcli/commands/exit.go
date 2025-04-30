package commands

import (
	"fmt"
	"os"
)

func CommandExit(config *Config, input []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
