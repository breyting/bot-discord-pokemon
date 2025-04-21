package commands

import (
	"fmt"
	"os"
)

type config struct {
	next     string
	previous string
}

func CommandExit(config *config, input []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
