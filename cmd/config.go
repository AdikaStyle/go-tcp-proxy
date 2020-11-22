package cmd

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Name      string `json:"name"`
	FromPort  int    `json:"from"`
	ToAddress string `json:"to"`
}

func ParseConfig(filePath string) (Config, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var out Config
	if err := json.Unmarshal(content, &out); err != nil {
		return Config{}, err
	}

	return out, nil
}
