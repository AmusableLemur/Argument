package main

import (
	"os"
	"strings"
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
	dbURI := "root:@tcp(127.0.0.1:3306)/argument"

	if os.Getenv("DB_PORT") != "" {
		// This is necessary to handle dynamic ports with Github actions
		strings.Replace(dbURI, "3306", os.Getenv("DB_PORT"), 1)
	}

	assert.Equal(t, dbURI, config.Database.URI)
}
