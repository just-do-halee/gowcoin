package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"

	"github.com/just-do-halee/gowcoin/tree/main/gow-node/cmn"
	"github.com/just-do-halee/gowcoin/tree/main/gow-node/db"
)

type Block struct {
	Data     []byte   `json:"data"`
	Hash     cmn.Hash `json:"hash"`
	PrevHash cmn.Hash `json:"prevHash,omitempty"`
	Height   uint64   `json:"height"`
}

func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	cmn.HandleError(encoder.Encode(b))
	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

func createBlock(data []byte, prevHash cmn.Hash, height uint64) *Block {
	block := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
		Height:   height,
	}

	payload := make([]byte, 0, len(data)+cmn.HashLength*2+8)

	heightBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(heightBytes, height)

	payload = append(payload, block.Data...)
	payload = append(payload, block.PrevHash.Bytes()...)
	payload = append(payload, heightBytes...)
	block.Hash = sha256.Sum256(payload)
	block.persist()
	return block
}
