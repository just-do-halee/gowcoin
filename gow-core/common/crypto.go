package common

import "crypto/rand"

const (
	// HashSize is the size of a hash in bytes.
	HashSize uint = 32
)

// Hash represents the [HashSize]byte array.
type Hash [HashSize]byte

// ToBytes converts the hash to a byte slice.
func (h *Hash) ToBytes() []byte {
	return h[:]
}

// NewRandomHash creates a new random hash.
func NewRandomHash() (Hash, error) {
	bytes := [HashSize]byte{}
	_, err := rand.Read(bytes[:])
	if err != nil {
		return bytes, err
	}

	return Hash(bytes), nil
}
