package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Hubert-Madej/slackmate/pkg/models"
)

// LoadConfig loads configuration from the given file.
func LoadConfig(configFileName string, config *models.Config, key []byte) error {
	// Try to open the config file
	file, err := os.Open(configFileName)
	if err != nil {
		return fmt.Errorf("❌ Missing config file: %v", err)
	}
	defer file.Close()

	// Decrypt file content
	encryptedData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("❌ Failed to read config file: %v", err)
	}

	decryptedData, err := Decrypt(encryptedData, key)
	if err != nil {
		return fmt.Errorf("❌ Failed to decrypt config file: %v", err)
	}

	// Decode config JSON
	err = json.Unmarshal(decryptedData, config)
	if err != nil {
		return fmt.Errorf("❌ Failed to decode config file: %v", err)
	}

	return nil
}

// SaveConfig saves configuration to the given file.
func SaveConfig(config models.Config, configFileName string, encryptionKey []byte) error {
	// Marshal config to JSON
	configJSON, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Encrypt config data
	encryptedData, err := Encrypt(configJSON, encryptionKey)
	if err != nil {
		return err
	}

	// Write encrypted data to file
	file, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(encryptedData)
	if err != nil {
		return err
	}

	return nil
}
