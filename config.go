package main

import (
	"io/ioutil"

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

// LoadConfig loads a config from a file
func LoadConfig(f string) Config {
	stream, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	return SetupConfig(stream)
}

// SetupConfig converts a byte stream to a Config object
func SetupConfig(stream []byte) Config {
	config := Config{}
	toml.Unmarshal(stream, &config)

	if "" == config.Title {
		config.Title = "Argument"
	}

	return config
}
