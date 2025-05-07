package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/breyting/pokedex-discord/pokedexcli/pokeapi"
	"github.com/bwmarrin/discordgo"
)

func CommandCatch(config *Config, data map[string]UserData, s *discordgo.Session, m *discordgo.MessageCreate, input ...string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("can not catch without a pokemon name")
	}

	pokemonInput := input[0]

	pokemonInfo, err := GetPokemon(config, pokemonInput)
	if err != nil {
		return "", err
	}

	return tryingCatch(pokemonInfo, data, s, m)
}

func tryingCatch(pokemonInfo pokeapi.Pokemon, data map[string]UserData, s *discordgo.Session, m *discordgo.MessageCreate) (string, error) {
	baseExperience := pokemonInfo.BaseExperience
	chance := rand.Intn(baseExperience)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Throwing a Pokeball at %s...\n", pokemonInfo.Name))
	if chance < 5 {
		// Shiny Pokemon
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s shiny was caught!\n", pokemonInfo.Name))

		new_pokemon := UserData{
			pokemonInfo,
			time.Now(),
			true,
		}

		sprite := pokemonInfo.Sprites.FrontShiny
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v\n", sprite))

		waitForNickname(s, m.Author.ID, m.ChannelID, new_pokemon, data)

		msg := "You may now inspect it with the inspect command.\n"
		return msg, nil
	} else if chance < 50 {
		// Normal Pokemon
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s was caught!\n", pokemonInfo.Name))
		new_pokemon := UserData{
			pokemonInfo,
			time.Now(),
			false,
		}

		sprite := pokemonInfo.Sprites.FrontDefault
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v\n", sprite))

		waitForNickname(s, m.Author.ID, m.ChannelID, new_pokemon, data)

		msg := "You may now inspect it with the inspect command.\n"

		return msg, nil
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s escaped!\n", pokemonInfo.Name))
		sprite := pokemonInfo.Sprites.BackDefault
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v\n", sprite))
		return "", nil
	}
}

func waitForNickname(s *discordgo.Session, userID, channelID string, newPokemon UserData, data map[string]UserData) {
	s.ChannelMessageSend(channelID, fmt.Sprintf("What nickname do you want to give to %s?\n", newPokemon.Pokemon.Name))
	timeout := time.After(30 * time.Second)
	messages := make(chan *discordgo.MessageCreate)

	// Temp handler for replies
	handler := func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == userID && m.ChannelID == channelID {
			messages <- m
		}
	}

	remove := s.AddHandler(handler)
	defer remove()

	select {
	case msg := <-messages:
		nickname := strings.TrimSpace(msg.Content)
		_, aleardyIn := data[nickname]
		if aleardyIn {
			s.ChannelMessageSend(channelID, fmt.Sprintf("❌ %s is already taken. Please choose another nickname.", nickname))
			waitForNickname(s, userID, channelID, newPokemon, data)
			return
		} else if strings.ToLower(nickname) != "no" && nickname != "" {
			data[nickname] = newPokemon
			SaveUserData(userID, data)
			s.ChannelMessageSend(channelID, fmt.Sprintf("✅ %s has been nicknamed **%s**!", newPokemon.Pokemon.Name, nickname))
		} else {
			nickname := getIncrementedName(newPokemon.Pokemon.Name, data)
			data[nickname] = newPokemon
			SaveUserData(userID, data)
			s.ChannelMessageSend(channelID, fmt.Sprintf("No nickname given for %s.", nickname))
		}

	case <-timeout:
		s.ChannelMessageSend(channelID, fmt.Sprintf("⏰ No response received. %s was saved without a nickname.", newPokemon.Pokemon.Name))
	}
}

func getIncrementedName(name string, data map[string]UserData) string {
	if _, exists := data[name]; exists {
		for i := 1; ; i++ {
			newName := fmt.Sprintf("%s_%d", name, i)
			if _, exists := data[newName]; !exists {
				return newName
			}
		}
	}
	return name
}
