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
	Next:          "https://pokeapi.co/api/v2/location-area/1",
}

func main() {
	commands.Init()

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
	args := strings.Fields(content)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "catch":
		reply, err := commands.CommandCatch(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "explore":
		reply, err := commands.CommandExplore(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "map":
		reply, err := commands.CommandMap(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "mapb":
		reply, err := commands.CommandMapb(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "pokedex":
		reply, err := commands.CommandPokedex(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "inspect":
		reply, err := commands.CommandInspect(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "help":
		reply, err := commands.CommandHelp(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	default:
		s.ChannelMessageSend(m.ChannelID, "Unknown command. Type `help` for a list of commands.")
	}

}

func sendReply(s *discordgo.Session, channelID, reply string, err error) {
	if err != nil {
		s.ChannelMessageSend(channelID, "❌ "+err.Error())
	} else {
		s.ChannelMessageSend(channelID, reply)
	}
}
