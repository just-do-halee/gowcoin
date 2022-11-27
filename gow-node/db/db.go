package db

import (
	"fmt"

	"github.com/just-do-halee/gowcoin/tree/main/gow-node/cmn"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName       = "gowcoin.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var DB *bolt.DB

func SaveBlock(hash cmn.Hash, data []byte) {
	fmt.Printf("Saving block %s\nData: %b", hash, data)
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		return bucket.Put(hash.Bytes(), data)
	})
	cmn.HandleError(err)
}

func SaveBlockchain(data []byte) {
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		return bucket.Put([]byte("checkpoint"), data)
	})
	cmn.HandleError(err)
}
func init() {
	var err error
	DB, err = bolt.Open(dbName, 0600, nil)
	cmn.HandleError(err)
	err = DB.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(dataBucket))
		cmn.HandleError(err)
		_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	cmn.HandleError(err)
}
