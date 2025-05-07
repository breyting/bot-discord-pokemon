package commands

import "github.com/bwmarrin/discordgo"

func CommandFree(config *Config, data map[string]UserData, s *discordgo.Session, m *discordgo.MessageCreate, input ...string) (string, error) {
	if len(input) == 0 {
		return "You need to specify the name of a pokemon to release!\n", nil
	}

	for key, _ := range data {
		if input[0] == key {
			delete(data, key)
			SaveUserData(m.Author.ID, data)
			return "You have released " + key + "!\n", nil
		}
	}
	return "You don't have " + input[0] + " in your Pokedex!\n", nil
}
