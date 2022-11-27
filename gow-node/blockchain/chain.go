package blockchain

import "github.com/just-do-halee/gowcoin/tree/main/gow-node/cmn"

type blockchain struct {
	NewestHash cmn.Hash `json:"newestHash"`
	Height     uint64   `json:"height"`
}

func (b *blockchain) AddBlock(data []byte) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height++
}

var Blockchain *blockchain

func init() {
	Blockchain = &blockchain{
		NewestHash: cmn.Hash{},
		Height:     0,
	}
	Blockchain.AddBlock([]byte("Genesis"))
}
