package core

import "testing"

func TestPostStoreContract(t *testing.T) {
	postStore := NewPostMemoryStore()
	threadStore := NewThreadMemoryStore()
	contract := NewPostStoreContract(postStore, threadStore)
	contract.Verify(t)
}
