package vault

import (
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/sdbox"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/version"
)

type Vaults []*Vault

func (v *Vaults) Len() int {
	return len(*v)
}

// Vault is a block in the blockchain.
// It contains a list of sdboxes.
// And header information.
type Vault struct {
	Info    Info           `json:"info"`
	Sdboxes *sdbox.Sdboxes `json:"sdboxes"`
}

// NewGenesis creates a new genesis vault with the given owner.
func NewGenesis(owner common.Hash) *Vault {
	return &Vault{
		Info: Info{
			Top: TopLevelInfo{
				Index:   NewIndex(),
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
func New(owner common.Hash, sdboxes *sdbox.Sdboxes, level Level, previousHash common.Hash, index Index) *Vault {
	return &Vault{
		Info: Info{
			Top: TopLevelInfo{
				Index:        index,
				Version:      version.CurrentVault,
				PreviousHash: previousHash,
				Level:        level,
			},
			Supervisor: SupervisorLevelInfo{
				Owner: owner,
			},
		},
		Sdboxes: sdboxes,
	}
}
