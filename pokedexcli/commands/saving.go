package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadUserData(userID string) (map[string]UserData, error) {
	filePath := fmt.Sprintf("data/%s.json", userID)

	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return map[string]UserData{}, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()
	var data map[string]UserData

	byteVal, err := io.ReadAll(file)
	json.Unmarshal(byteVal, &data)
	return data, err
}

func SaveUserData(userID string, data map[string]UserData) error {
	filePath := fmt.Sprintf("data/%s.json", userID)
	os.MkdirAll("data", 0755)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}
