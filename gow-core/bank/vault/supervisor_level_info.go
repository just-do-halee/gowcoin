package vault

import "github.com/just-do-halee/gowcoin/tree/main/gow-core/common"

// The timestamp of the vault locked by the supervisor.
type LockedTime uint64

// It is a Nonce that is used to lock a vault.
type Password uint64

// SupervisorLevelInfo is the information of a vault that
// can be changed by the supervisor.
type SupervisorLevelInfo struct {
	Owner      common.Hash `json:"owner"`
	LockedTime LockedTime  `json:"locked_time"`
	Password   Password    `json:"password"`
}
