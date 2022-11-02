package vault

import (
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/version"
)

// The difficulty of locking a vault.
type Level uint64

const InitialLevel Level = 0

// TopLevelInfo is the top level information of a vault that
// cannot be changed after the vault is created.
type TopLevelInfo struct {
	Version      version.Vault
	PreviousHash common.Hash
	Level        Level
}
