package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// PreloadEncryptionKey generates a new encryption key if not provided or decodes the provided key.
func PreloadEncryptionKey(k string) ([]byte, error) {
	if k == "" {
		fmt.Println("To safely store your configuration data locally, SlackMate will generate an encryption key.")
		fmt.Println("WARNING: Keep this key private and do not share it with anyone!")
		fmt.Println("Press any key to display your encryption key")
		fmt.Scanln()
		// Generate a new encryption key if not provided
		key := make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			return nil, err
		}
		fmt.Printf("AES Key: %x\n", key)
		fmt.Println("Encryption key will be used during the next encryption operation and then will be discarded.")
		return key, nil
	}

	// Decode encryption key if provided
	decoded, err := hex.DecodeString(k)
	if err != nil {
		fmt.Println("Error decoding hexadecimal string:", err)
		return nil, fmt.Errorf("error decoding encryption key: %v", err)
	}

	var keyBytes []byte
	if len(decoded) >= 32 {
		keyBytes = decoded[:32] // Truncate if longer than 32 bytes
	} else {
		keyBytes = make([]byte, 32)
		copy(keyBytes, decoded) // Pad if shorter than 32 bytes
	}

	return keyBytes, nil
}



func Encrypt(data []byte, encryptionKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedData := gcm.Seal(nonce, nonce, data, nil)
	return encryptedData, nil
}

func Decrypt(data []byte, encryptionKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("invalid encrypted data")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}