package vault

import "github.com/just-do-halee/gowcoin/tree/main/gow-core/common"

// StateLevelInfo is the information of a bank that
// can be changed by the state.
type StateLevelInfo struct {
	StateRootHash common.Hash `json:"state_root_hash"`
	SdboxRootHash common.Hash `json:"sdbox_root_hash"`
}
