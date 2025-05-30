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
	PokeapiClient: pokeClient,
	Next:          pokeapi.BaseURL + "/location-area/1",
}
var welcomedUsers = make(map[string]bool)

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

	userID := m.Author.ID
	data, err := commands.LoadUserData(userID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Erreur chargement utilisateur."+err.Error())
		return
	}

	content := strings.ToLower(m.Content)
	args := strings.Fields(content)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "hi", "hello", "hey":
		sendWelcomeDM(s, m.Author.ID, welcomedUsers[m.Author.ID])
		welcomedUsers[m.Author.ID] = true

	case "catch":
		reply, err := commands.CommandCatch(cfg, data, s, m, args[1:]...)
		if err == nil {
			commands.SaveUserData(userID, data)
		}
		sendReply(s, m.ChannelID, reply, err)

	case "explore":
		reply, err := commands.CommandExplore(cfg, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "map":
		reply, err := commands.CommandMap(cfg, args[1:]...)
		s.ChannelMessageSend(m.ChannelID, "Trying to get the next locations...")
		sendReply(s, m.ChannelID, reply, err)

	case "mapb":
		reply, err := commands.CommandMapb(cfg, args[1:]...)
		s.ChannelMessageSend(m.ChannelID, "Trying to get the previous locations...")
		sendReply(s, m.ChannelID, reply, err)

	case "pokedex":
		reply, err := commands.CommandPokedex(cfg, data, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "inspect":
		reply, err := commands.CommandInspect(cfg, data, s, m, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "help":
		reply, err := commands.CommandHelp(args[1:]...)
		sendReply(s, m.ChannelID, reply, err)

	case "free":
		reply, err := commands.CommandFree(cfg, data, s, m, args[1:]...)
		sendReply(s, m.ChannelID, reply, err)
	}

}

func sendReply(s *discordgo.Session, channelID, reply string, err error) {
	if err != nil {
		s.ChannelMessageSend(channelID, "❌ "+err.Error())
	} else {
		s.ChannelMessageSend(channelID, reply)
	}
}

func sendWelcomeDM(s *discordgo.Session, userID string, welcomed bool) {
	var msg string
	if welcomed {
		msg += "I already welcomed you... Whatever...\n\n"
	}

	channel, err := s.UserChannelCreate(userID)
	if err != nil {
		fmt.Println("Error creating DM channel:", err)
		return
	}

	msg += "**👋 Welcome to the Pokedex Bot!**\n\n" +
		"Here, you can explore the Pokemon world, catch Pokemon, and build your own Pokedex — all within Discord!\n\n" +
		"**Basic commands:**\n" +
		"`help` – List all available commands\n" +
		"`map` and `mapb` – Display the 20 next or previous locations of the Pokemon world\n" +
		"`explore [location]` – Displays all the Pokemon available in the location\n" +
		"`catch [pokemon]` – Try to catch a Pokémon\n" +
		"`inspect [pokemon]` – See details of a captured Pokemon\n" +
		"`pokedex` – View your Pokedex with your catched Pokemons\n\n" +
		"Good luck, Trainer! 🔍🎒"

	s.ChannelMessageSend(channel.ID, msg)
}
