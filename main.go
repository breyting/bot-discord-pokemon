package main

import (
	"time"

	"github.com/breyting/pokedex-discord/pokedexcli/commands"
	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &commands.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
