package core

type txn interface {
	GetDoc()
	InsertDoc()
	UpdateDoc()
	DeleteDoc()
}

type txnImpl struct{

}

func (t *txnImpl) GetDoc(key []byte, collection []byte) []byte{
	return nil
}

func (d *Doc) InsertDoc(key []byte, value []byte, table string) []byte{
	return nil
}


func (d *Doc) UpdateDoc(key []byte, value []byte, table string) []byte{
	return nil
}

func (d *Doc) DeleteDoc(key []byte, value []byte, table string) []byte{
	return nil
}








