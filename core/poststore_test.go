package core

import "testing"

func TestCreateAndReadPost(t *testing.T) {
	store := NewPostMemoryStore()
	contract := NewPostStoreContract(store)
	contract.CreateAndReadPostContract(t)
}

func TestCreateAndReadMultiplePosts(t *testing.T) {
	store := NewPostMemoryStore()
	contract := NewPostStoreContract(store)
	contract.CreateAndReadMultiplePostsContract(t)
}

func TestReadPostReturnsErrorWhenIdDoesNotExist(t *testing.T) {
	store := NewPostMemoryStore()
	contract := NewPostStoreContract(store)
	contract.ReadPostReturnsErrorWhenIdDoesNotExistContract(t)
}

func TestReadByThreadId(t *testing.T) {
	store := NewPostMemoryStore()
	contract := NewPostStoreContract(store)
	contract.ReadByThreadIdContract(t)
}
