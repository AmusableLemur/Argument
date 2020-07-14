package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := SetupConfig([]byte(`
        [Database]
        Username = "root"
        Password = "secret"`),
	)

	assert.Equal(t, config.Database.Username, "root")
	assert.Equal(t, config.Database.Password, "secret")
	assert.Equal(t, config.Title, "Argument")
}

func TestLoadingConfig(t *testing.T) {
	config, _ := LoadConfig("config-sample.toml")

	assert.Equal(t, "Argument", config.Title)
}

func TestLoadingNonexistentConfig(t *testing.T) {
	_, err := LoadConfig("non-existent-config.toml")

	if err == nil {
		t.Errorf("Missing config file should throw an error")
	}
}

func TestDatabaseURIConfig(t *testing.T) {
	config, _ := LoadConfig("config-sample.toml")

	assert.Equal(t, "root:@tcp(127.0.0.1:3306)/argument", config.Database.URI)
}
