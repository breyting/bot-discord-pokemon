package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/breyting/pokedex-discord/pokedexcli/commands"
	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
	"github.com/bwmarrin/discordgo"
)

var pokeClient = pokeapi.NewClient(5*time.Second, 5*time.Minute)
var cfg = &commands.Config{
	CaughtPokemon: map[string]pokeapi.Pokemon{},
	PokeapiClient: pokeClient,
}

func main() {

	//startRepl(cfg)

	token := Bot_token
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	content := strings.ToLower(m.Content)
	if strings.HasPrefix(content, "!catch ") {
		args := strings.Fields(content)[1:]
		result, err := commands.CommandCatch(cfg, args)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Erreur : "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, result)
	}
}
