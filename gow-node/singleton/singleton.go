package singleton

import (
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/bank"
	"github.com/just-do-halee/gowcoin/tree/main/gow-core/common"
)

var k *common.Hash

func GetKey() *common.Hash {
	return k
}

var b *bank.Bank

func GetBank() *bank.Bank {
	return b
}

func init() {
	key, err := common.NewRandomHash()
	if err != nil {
		panic(err)
	}
	k = &key

	b = bank.New(key)
}
