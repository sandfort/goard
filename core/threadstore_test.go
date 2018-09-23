package core

import "testing"

func TestThreadStoreContract(t *testing.T) {
	store := NewThreadMemoryStore()
	contract := NewThreadStoreContract(store)
	contract.Verify(t)
}
