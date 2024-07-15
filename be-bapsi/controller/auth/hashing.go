package auth

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a salted hash for the input password
func HashPassword(password string) string {
	// Generate a random salt
	salt, err := generateSalt()
	if err != nil {
		return ""
	}

	// Hash the password with the generated salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	// Combine salt and hashed password
	saltAndHash := append([]byte(salt), hashedPassword...)

	// Encode the combined salt and hashed password to a base64 string
	encodedHash := base64.StdEncoding.EncodeToString(saltAndHash)
	return encodedHash
}

// CheckPasswordHash checks if the password matches the hash
func CheckPasswordHash(password, encodedHash string) bool {
	// Decode the base64 encoded hash
	saltAndHash, err := base64.StdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false
	}

	// Extract the salt (first 16 bytes)
	if len(saltAndHash) < 16 {
		return false
	}
	salt := saltAndHash[:16]
	hashedPassword := saltAndHash[16:]

	// Hash the password with the extracted salt and compare it to the stored hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password+string(salt)))
	return err == nil
}

// generateSalt generates a random salt
func generateSalt() (string, error) {
	// Generate 16 random bytes
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Return the salt as a string
	return string(salt), nil
}
