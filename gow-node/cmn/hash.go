package cmn

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// HashLength is the length of a hash in bytes
const HashLength = 32

type Hash [HashLength]byte

func (h Hash) Bytes() []byte { return h[:] }

func (h Hash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }

func (h Hash) Hex() string { return hex.EncodeToString(h[:]) }

func (h Hash) String() string { return h.Hex() }

func NewRandomHash() (Hash, error) {
	bytes := [HashLength]byte{}
	_, err := rand.Read(bytes[:])
	return bytes, err
}
