package core


type collection interface {
	createCollection()
	DeleteCollection()
	UpdateCollection()
	GetCollectionInfo()
}

type collectionImpl struct {

}

