package core

import (
	badger "github.com/dgraph-io/badger/v2"
	"log"
)

type db interface {
	InitDB()
	GetDB()
}

type dbImpl struct {
	db *badger.DB
	dbPath string
}


func (dbImpl *dbImpl)dbInit(){
	var err error
	dbImpl.db, err = badger.Open(badger.DefaultOptions(dbImpl.dbPath))
	if err != nil {
		log.Fatal(err)
	}
}

func (dbImpl *dbImpl) GetDB() *badger.DB{
	return dbImpl.db
}
