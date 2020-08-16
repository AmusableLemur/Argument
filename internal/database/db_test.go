package database

import (
	"testing"

	"github.com/AmusableLemur/Argument/internal/config"
	"github.com/stretchr/testify/assert"
)

func setupTest() {
	Disconnect()
	Connect(config.Conf.Database.URI)
}

func TestBrokenDB(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("If database is unreachable we should crash")
		}
	}()

	Disconnect()
	Connect("root:asdf@tcp(127.0.0.1:12345)/argument")
}

func TestConnectingTwice(t *testing.T) {
	Connect(config.Conf.Database.URI)
	Connect(config.Conf.Database.URI)
}

func TestGetPosts(t *testing.T) {
	setupTest()
	assert.NotNil(t, GetPosts())
}
