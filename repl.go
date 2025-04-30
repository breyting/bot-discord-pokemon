package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	commands "github.com/breyting/pokedex-discord/pokedexcli/commands"
	pokecache "github.com/breyting/pokedex-discord/pokedexcli/pokecache"
)

var cache = pokecache.NewCache((5 * time.Second))

func startRepl(conf *commands.Config) {
	commands.Init()
	conf.Next = "https://pokeapi.co/api/v2/location-area/1"

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scan.Scan()
		text := scan.Text()
		input := cleanInput(text)

		if len(input) == 0 {
			continue
		}

		command, exists := commands.ListOfCommands[input[0]]
		if exists {
			err := command.Callback(conf, input[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
