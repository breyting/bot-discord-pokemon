package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

func CommandMap(config *Config, input ...string) (string, error) {
	msg := "Here are the 20 next locations:\n"

	for i := 0; i < 20; i++ {
		id := path.Base(config.Next)

		cache := config.PokeapiClient.Cache
		val, ok := cache.Get(id)
		if ok {
			msg += string(val) + "\n"
		} else {
			res, err := http.Get(config.Next)
			if err != nil {
				return "", fmt.Errorf("Error fetching data: %s", err)
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
		NextId, _ := strconv.Atoi(id)
		NextId += 1
		config.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d", NextId)
	}
	return msg, nil
}
