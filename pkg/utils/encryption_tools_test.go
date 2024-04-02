package utils

import (
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	originalData := []byte("SlackMate is cool.")

	encryptionKey, err := PreloadEncryptionKey("")
	if err != nil {
		t.Fatalf("Failed to generate encryption key: %v", err)
	}

	encryptedData, err := Encrypt(originalData, encryptionKey)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decryptedData, err := Decrypt(encryptedData, encryptionKey)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Ensure original data matches decrypted data
	if string(originalData) != string(decryptedData) {
		t.Errorf("Decrypted data does not match original data")
	}
}
