package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Function to generate a fixed IV
func generateIV() []byte {
	iv := make([]byte, aes.BlockSize)
	return iv
}

func SecretKey() []byte {
	secretKey := "FaizalZametsy"
	key := []byte(secretKey)

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		hash := sha256.Sum256(key)
		key = hash[:]
	}

	return key
}

// Function to encrypt string with AES
func Encrypt(text string) string {
	key := SecretKey()

	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("ENCRYPT: %s", err.Error())
		return ""
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := generateIV() // Generate a fixed IV
	copy(ciphertext[:aes.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)[21:]
}

// Function to decrypt string encrypted with AES
func Decrypt(ciphertext string) string {
	key := SecretKey()

	decodedCiphertext, err := base64.URLEncoding.DecodeString("AAAAAAAAAAAAAAAAAAAAA" + ciphertext)
	if err != nil {
		fmt.Printf("DECRYPT: %s", err.Error())
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("DECRYPT: %s", err.Error())
		return ""
	}

	if len(decodedCiphertext) < aes.BlockSize {
		fmt.Printf("error: ciphertext too short")
		return ""
	}
	iv := decodedCiphertext[:aes.BlockSize]
	decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decodedCiphertext, decodedCiphertext)

	return string(decodedCiphertext)
}
