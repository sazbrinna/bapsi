package id

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func NewV6() uuid.UUID {
	uuidBytes := make([]byte, 16)

	// Set version (bits 12-15) to 6
	uuidBytes[6] = (6 << 4) | (uuidBytes[6] & 0x0f)

	// Set variant (bits 16-17) to 10xx
	uuidBytes[8] = (uuidBytes[8] & 0xbf) | 0x80

	// Fill with random data
	rand.Seed(time.Now().UnixNano())
	rand.Read(uuidBytes)

	return uuid.UUID(uuidBytes)
}
