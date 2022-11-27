package single

import (
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
	"github.com/just-do-halee/gowcoin/tree/main/gow-node/cmn"
)

var Key *cmn.Hash

var Bank *bank.Bank

func init() {
	key, err := cmn.NewRandomHash()
	if err != nil {
		panic(err)
	}
	Key = &key

	Bank = bank.New(common.Hash(*Key))
}
