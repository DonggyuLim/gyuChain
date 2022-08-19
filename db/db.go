package db

import (
	"log"
	"sync"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/gyuChain/utils"
)

//Our Blockchain is used badgerDB

type ChainDB struct {
	DB    *badger.DB
	mutex sync.Mutex
}

var chainDB *ChainDB
var once sync.Once

func InitChainDB() {
	if chainDB == nil {
		once.Do(func() {
			dbPointer, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
			utils.PanicErr(err)
			chainDB.DB = dbPointer
		})
	} else {
		return
	}
}

func ChainDBClose() {
	err := chainDB.DB.Close()
	if err == nil {
		log.Printf("database closed %s", err)
	} else {
		log.Println("failed to close database", err)
	}
}
