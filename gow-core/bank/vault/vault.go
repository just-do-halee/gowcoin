package vault

import (
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/version"
)

// Vault is a block in the blockchain.
// It contains a list of sdboxes.
// And header information.
type Vault struct {
	info    Info
	sdboxes []*Sdbox
}

// NewGenesis creates a new genesis vault with the given owner.
func NewGenesis(owner common.Hash) *Vault {
	return &Vault{
		info: Info{
			Top: TopLevelInfo{
				Version: version.CurrentVault,
				Level:   InitialLevel,
			},
			Supervisor: SupervisorLevelInfo{
				Owner: owner,
			},
		},
	}
}

// New creates a new vault with the given owner and sdboxes, level, and previous hash.
func New(owner common.Hash, sdboxes []*Sdbox, level Level, previousHash common.Hash) *Vault {
	return &Vault{
		info: Info{
			Top: TopLevelInfo{
				Version:      version.CurrentVault,
				PreviousHash: previousHash,
				Level:        level,
			},
			Supervisor: SupervisorLevelInfo{
				Owner: owner,
			},
		},
		sdboxes: sdboxes,
	}
}
