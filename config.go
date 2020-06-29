package main

import (
	toml "github.com/pelletier/go-toml"
)

// Config specifies the availabel configuration values
type Config struct {
	Title    string
	Database struct {
		Username string
		Password string
	}
}

// LoadConfig converts a byte stream to a Config object
func LoadConfig(stream []byte) Config {
	config := Config{}
	toml.Unmarshal(stream, &config)

	if "" == config.Title {
		config.Title = "Argument"
	}

	return config
}
