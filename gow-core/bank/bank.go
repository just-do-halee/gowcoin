package bank

import (
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/vault"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
)

// Bank is a blockchain that stores vaults.
type Bank struct {
	vaults []*vault.Vault
}

// New creates a new bank with a genesis vault.
func New(owner common.Hash) *Bank {
	return &Bank{
		[]*vault.Vault{vault.NewGenesis(owner)},
	}
}
