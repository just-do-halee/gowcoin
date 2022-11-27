package vault

import (
	"math/big"

	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/version"
)

// The difficulty of locking a vault.
type Level = uint64

// The index of a vault.
type Index = *big.Int

func NewIndex() Index {
	return new(big.Int)
}

const InitialLevel Level = 0

// TopLevelInfo is the top level information of a vault that
// cannot be changed after the vault is created.
type TopLevelInfo struct {
	Index        Index         `json:"index"`
	Version      version.Vault `json:"version"`
	PreviousHash common.Hash   `json:"previous_hash,omitempty"`
	Level        Level         `json:"level"`
}
