package core

import "testing"

func TestContract(t *testing.T) {
	store := NewPostMemoryStore()
	contract := NewPostStoreContract(store)
	contract.Verify(t)
}
