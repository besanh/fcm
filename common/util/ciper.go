package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var secretKey = []byte("anhle@!*2025#%0~")

func Encrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// aesGCM.Seal => create ciphertext
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	// Convert to base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedText string) (bool, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return false, err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return false, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, fmt.Errorf("invalid ciphertext")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, err
	}

	return string(plaintext) == "anhle@!*2025#%0~", nil
}
