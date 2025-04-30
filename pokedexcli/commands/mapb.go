package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

func CommandMapb(config *Config, input ...string) (string, error) {
	id := path.Base(config.Next)
	nextId, _ := strconv.Atoi(id)
	msg := "Here are the 20 previous locations:\n"

	if nextId < 40 {
		return "", errors.New("can not show previous locations if not at least 40 locations are shown")
	}

	nextId -= 40
	config.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", nextId)
	config.Previous = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", nextId-1)

	for i := 0; i < 20; i++ {
		id := path.Base(config.Next)
		cache := config.PokeapiClient.Cache
		val, ok := cache.Get(id)
		if ok {
			msg += string(val) + "\n"
		} else {
			res, err := http.Get(config.Next)
			if err != nil {
				return "", err
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				return "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				return "", err
			}

			location := map[string]string{}
			json.Unmarshal(body, &location)

			msg += location["name"] + "\n"
			cache.Add(id, []byte(location["name"]))
		}

		config.Previous = config.Next
		id = path.Base(config.Previous)
		nextId, _ := strconv.Atoi(id)
		nextId += 1
		config.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", nextId)
	}
	return msg, nil
}
