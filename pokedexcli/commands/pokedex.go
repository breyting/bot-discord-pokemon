package commands

import "fmt"

func CommandPokedex(config *config, input []string) error {
	if len(ownPokedex) == 0 {
		return fmt.Errorf("you didn't catch any pokemon yet")
	}
	fmt.Println("Your Pokedex :")
	for _, val := range ownPokedex {
		fmt.Printf("- %s\n", val.Name)
	}
	return nil
}
