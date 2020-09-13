package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	setupTest()
	assert.NotNil(t, GetPosts())
}
