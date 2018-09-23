package core

import "testing"

func TestPostStoreContract(t *testing.T) {
	store := NewPostMemoryStore()
	contract := NewPostStoreContract(store)
	contract.Verify(t)
}
