package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := LoadConfig([]byte(`
        [Database]
        Username = "root"
        Password = "secret"`),
	)

	assert.Equal(t, config.Database.Username, "root")
	assert.Equal(t, config.Database.Password, "secret")
	assert.Equal(t, config.Title, "Argument")
}
