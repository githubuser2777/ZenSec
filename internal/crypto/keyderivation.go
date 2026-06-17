package crypto

import (
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

const (
	saltSize    = 16
	keySize     = 32
	argonTime   = 3
	argonMemory = 64 * 1024
	argonThreads= 4
)

// DeriveKey derives a 32-byte key from a password and a salt using Argon2id.
func DeriveKey(password []byte, salt []byte) []byte {
	return argon2.IDKey(password, salt, argonTime, argonMemory, uint8(argonThreads), keySize)
}

// GenerateSalt generates a random salt of a specified size (default 16 bytes).
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
