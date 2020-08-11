package config

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
	config, _ := LoadConfig("../../config-sample.toml")

	assert.Equal(t, "Argument", config.Title)
}

func TestLoadingNonexistentConfig(t *testing.T) {
	_, err := LoadConfig("non-existent-config.toml")

	if err == nil {
		t.Errorf("Missing config file should throw an error")
	}
}

func TestDatabaseURIConfig(t *testing.T) {
	config, _ := LoadConfig("../../config-sample.toml")
	dbURI := "root:@tcp(127.0.0.1:3306)/argument"

	if "" != os.Getenv("DB_PORT") {
		// This is necessary to handle dynamic ports with Github actions
		dbURI = strings.Replace(dbURI, "3306", os.Getenv("DB_PORT"), 1)
	}

	assert.Equal(t, dbURI, config.Database.URI)
}

func TestEnvironmentConfig(t *testing.T) {
	originalHost := os.Getenv("DB_HOST")
	originalPort := os.Getenv("DB_PORT")

	os.Setenv("DB_HOST", "127.0.0.2")
	os.Setenv("DB_PORT", "1234")

	config, _ := LoadConfig("../../config-sample.toml")

	assert.Equal(t, "root:@tcp(127.0.0.2:1234)/argument", config.Database.URI)

	// Restore the original values for the later tests
	os.Setenv("DB_HOST", originalHost)
	os.Setenv("DB_PORT", originalPort)
}
