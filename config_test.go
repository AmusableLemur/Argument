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
	config := LoadConfig("config-sample.toml")

	assert.Equal(t, "Argument", config.Title)
}

func TestLoadingNonexistentConfig(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Missing config file should throw an error")
		}
	}()

	LoadConfig("non-existent-config.toml")
}

func TestDatabaseURIConfig(t *testing.T) {
	config := LoadConfig("config-sample.toml")

	assert.Equal(t, "root:@tcp(127.0.0.1:3306)/argument", config.Database.URI)
}
