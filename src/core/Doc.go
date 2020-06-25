package core


type Doc struct {
	key       []byte
	value       []byte
	expiresAt uint64
	version   uint64
	err       error
}
