package bank

import (
	"fmt"

	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/sdbox"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank/vault"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
)

// Bank is a blockchain that stores vaults.
type Bank struct {
	owner  common.Hash
	vaults vault.Vaults
}

// New creates a new bank with a genesis vault.
func New(owner common.Hash) *Bank {
	return &Bank{
		owner,
		vault.Vaults{vault.NewGenesis(owner)},
	}
}

// AddOwnerVault adds a new vault to the bank with the owner.
func (b *Bank) AddOwnerVault(sdboxes *sdbox.Sdboxes) {
	index := vault.NewIndex().SetInt64(int64(b.vaults.Len()))
	vault := vault.New(b.owner, sdboxes, 0, common.Hash{}, index)
	b.vaults = append(b.vaults, vault)
}

// AddVault adds a new vault to the bank.
func (b *Bank) AddVault(vault *vault.Vault) {
	b.vaults = append(b.vaults, vault)
}

func (b *Bank) AllVaults() *vault.Vaults {
	return &b.vaults
}

type ErrVaultNotFound struct {
	Index *vault.Index
}

func (e ErrVaultNotFound) Error() string {
	return fmt.Sprintf("Index: %d, the vault was not found", *e.Index)
}

func (b *Bank) GetVault(index vault.Index) (*vault.Vault, error) {
	if int(index.Int64()) >= b.vaults.Len() {
		return nil, ErrVaultNotFound{&index}
	}
	return b.vaults[index.Int64()], nil
}
