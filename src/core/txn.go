package core

import "github.com/dgraph-io/badger/v2"

type txn interface {
	GetDoc()
	InsertDoc()
	UpdateDoc()
	DeleteDoc()
}

type txnImpl struct{
	d dbImpl
	doc Doc
}

func GetTxn() txnImpl{
	var t txnImpl
	return t
}

func (t *txnImpl) GetDoc(key []byte, collection []byte) *badger.Item{
	db := t.d.GetDB()
	txn := db.NewTransaction(true)
	item, e  := txn.Get(key)
	if(e != nil) {
		txn.Discard()
	}
	_ = txn.Commit()
	return item
}

func (t *txnImpl) InsertDoc(key []byte, value []byte, collection []byte) error{
	db := t.d.GetDB()
	txn := db.NewTransaction(true)
	e := badger.NewEntry(key, value)
	err  := txn.SetEntry(e)
	if(err != nil){
		txn.Discard()
		return err
	}
	_ = txn.Commit()
	return nil
}


func (t *txnImpl) UpdateDoc(key []byte, value []byte, collection []byte) error{
	db := t.d.GetDB()
	txn := db.NewTransaction(true)
	err  := txn.Set(key, value)
	if(err != nil) {
		txn.Discard()
		return err
	}
	_ = txn.Commit()
	return nil
}

func (t *txnImpl) DeleteDoc(key []byte, collection []byte) error{
	db := t.d.GetDB()
	txn := db.NewTransaction(true)
	err  := txn.Delete(key)
	if(err != nil) {
		txn.Discard()
		return err
	}
	_ = txn.Commit()
	return nil
	return nil
}








