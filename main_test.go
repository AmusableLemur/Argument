package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "New Post")
}

func TestConfigSetup(t *testing.T) {
	config := setupConfig("config-sample.toml")

	assert.Equal(t, "Argument", config.Title)
}

func TestFailedConfigSetup(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Missing config file should throw an error")
		}
	}()

	setupConfig("non-existent-config.toml")
}
