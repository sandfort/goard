package core

import (
	"testing"
)

func TestCreateAndReadThread(t *testing.T) {
	store := NewThreadMemoryStore()
	id := store.CreateThread(Thread{
		Title: "yo",
	})
	result, _ := store.ReadThread(id)

	if result.Title != "yo" {
		t.Fail()
	}
}

func TestCreateAndReadMultipleThreads(t *testing.T) {
	store := NewThreadMemoryStore()
	id1 := store.CreateThread(Thread{Title: "first"})
	id2 := store.CreateThread(Thread{Title: "second"})
	p1, _ := store.ReadThread(id1)
	p2, _ := store.ReadThread(id2)

	if p1.Title != "first" {
		t.Fail()
	}

	if p2.Title != "second" {
		t.Fail()
	}
}

func TestReadReturnsErrorWhenIdDoesNotExist(t *testing.T) {
	store := NewThreadMemoryStore()
	_, err := store.ReadThread(1)

	if err == nil {
		t.Fail()
	}
}

func TestCreateAndReadAllThreads(t *testing.T) {
	store := NewThreadMemoryStore()
	store.CreateThread(Thread{Title: "first"})
	store.CreateThread(Thread{Title: "second"})
	threads := store.ReadAllThreads()

	if len(threads) != 2 {
		t.Errorf("Expected 2 but got %d", len(threads))
	}
}
