package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CommandInspect(config *Config, data map[string]UserData, s *discordgo.Session, m *discordgo.MessageCreate, input ...string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("you need to specify a pokemon name")
	}

	pokemon := input[0]
	_, ok := data[pokemon]
	if ok {
		pokemonInfo := data[pokemon].Pokemon
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Name: %s\n", pokemonInfo.Name))

		if data[pokemon].IsShiny {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s is a shiny Pokemon !\n", pokemon))
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s\n", pokemonInfo.Sprites.FrontShiny))
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s\n", pokemonInfo.Sprites.FrontDefault))
		}

		msg := fmt.Sprintf("Height: %d\n", pokemonInfo.Height)
		msg += fmt.Sprintf("Weight: %d\n", pokemonInfo.Weight)
		msg += "Stats:\n"
		for _, val := range pokemonInfo.Stats {
			msg += fmt.Sprintf("-%s: %d\n", val.Stat.Name, val.BaseStat)
		}
		msg += "Types:\n"
		for _, val := range pokemonInfo.Types {
			msg += fmt.Sprintf("- %s\n", val.Type.Name)
		}
		return msg, nil
	} else {
		return "", fmt.Errorf("you have not caught that pokemon")
	}
}
