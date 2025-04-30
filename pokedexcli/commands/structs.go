package commands

import (
	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

type Config struct {
	PokeapiClient pokeapi.Client
	Next          string
	Previous      string
	CaughtPokemon map[string]pokeapi.Pokemon
}
