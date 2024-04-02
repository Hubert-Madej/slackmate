package utils

import (
	"encoding/hex"
	"os"
	"testing"

	"github.com/Hubert-Madej/slackmate/pkg/models"
)

var TEST_TOKEN = "604421cbbf18535c06b5a583475e123995eeb39c1939a1c18ca62b733b166f31"

func TestConfigIO(t *testing.T) {
	key, err := hex.DecodeString(TEST_TOKEN)
	if err != nil {
		t.Fatalf("Failed to decode test token: %v", err)
	}
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "config_test.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Sample config data
	expectedConfig := models.Config{
		APIToken:       "secret-token",
		DefaultChannel: "channel-1",
	}

	// Save config to temporary file
	if err := SaveConfig(expectedConfig, tempFile.Name(), []byte(key)); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Load config from temporary file
	var actualConfig models.Config
	if err := LoadConfig(tempFile.Name(), &actualConfig, []byte(key)); err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Compare expected and actual config
	if !configEqual(expectedConfig, actualConfig) {
		t.Errorf("Loaded config does not match expected config")
	}
}

func configEqual(a, b models.Config) bool {
	if a.APIToken != b.APIToken {
		return false
	}

	if a.DefaultChannel != b.DefaultChannel {
		return false
	}

	return true
}
