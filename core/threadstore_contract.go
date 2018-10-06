package core

import "testing"

func NewThreadStoreContract(store ThreadStore) *threadStoreFixture {
	return &threadStoreFixture{store: store}
}

type threadStoreFixture struct {
	store ThreadStore
}

type ThreadStoreContract interface {
	Verify(t *testing.T)
	createAndReadThread(t *testing.T)
	createAndReadAllThreads(t *testing.T)
	createAndReadMultipleThreads(t *testing.T)
	readReturnsErrorWhenIdDoesNotExist(t *testing.T)
}

func (f *threadStoreFixture) Verify(t *testing.T) {
	f.createAndReadThread(t)
	f.createAndReadAllThreads(t)
	f.createAndReadMultipleThreads(t)
	f.readReturnsErrorWhenIdDoesNotExist(t)
}

func (f *threadStoreFixture) createAndReadThread(t *testing.T) {
	id := f.store.CreateThread(Thread{
		Title: "yo",
	})
	result, _ := f.store.ReadThread(id)

	if result.Title != "yo" {
		t.Errorf("Expected title to be \"yo\" but was %s", result.Title)
	}

	f.store.DeleteThread(id)
}

func (f *threadStoreFixture) createAndReadMultipleThreads(t *testing.T) {
	id1 := f.store.CreateThread(Thread{Title: "first"})
	id2 := f.store.CreateThread(Thread{Title: "second"})
	p1, _ := f.store.ReadThread(id1)
	p2, _ := f.store.ReadThread(id2)

	if p1.Title != "first" {
		t.Errorf("Expected title to be \"first\" but was %s", p1.Title)
	}

	if p2.Title != "second" {
		t.Errorf("Expected title to be \"second\" but was %s", p2.Title)
	}

	f.store.DeleteThread(id1)
	f.store.DeleteThread(id2)
}

func (f *threadStoreFixture) readReturnsErrorWhenIdDoesNotExist(t *testing.T) {
	_, err := f.store.ReadThread(1)

	if err == nil {
		t.Error("Expected error")
	}
}

func (f *threadStoreFixture) createAndReadAllThreads(t *testing.T) {
	id1 := f.store.CreateThread(Thread{Title: "first"})
	id2 := f.store.CreateThread(Thread{Title: "second"})
	threads := f.store.ReadAllThreads()

	if len(threads) != 2 {
		t.Errorf("Expected 2 but got %d", len(threads))
	}

	f.store.DeleteThread(id1)
	f.store.DeleteThread(id2)
}

func (f *threadStoreFixture) createAndDeleteThread(t *testing.T) {
	id := f.store.CreateThread(Thread{Title: "thread"})

	f.store.DeleteThread(id)

	_, err := f.store.ReadThread(id)

	if err == nil {
		t.Error("Expected to return error")
	}
}
