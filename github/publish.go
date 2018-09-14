package github

import (
	"encoding/json"
	"os"
)

type config struct {
	GithubUsername string `json:"github_username"`
	GithubPassword string `json:"github_password"`
}

type github struct {
	AppConfig config
}

func New(file string) *github {
	conf, _ := loadConfiguration(file)

	return &github{
		AppConfig: conf,
	}
}

func loadConfiguration(file string) (config, error) {
	var config config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}

func (g github) Publish(lines []string) error {
	return nil
}