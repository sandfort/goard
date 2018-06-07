package mem

import (
	"testing"

	"github.com/sandfort/goard/core"
)

func TestCreateAndReadThread(t *testing.T) {
	store := NewThreadStore()
	id := store.CreateThread(core.Thread{
		Title: "yo",
		Body:  "the body of the thread",
	})
	result, _ := store.ReadThread(id)

	if result.Title != "yo" || result.Body != "the body of the thread" {
		t.Fail()
	}
}

func TestCreateAndReadMultipleThreads(t *testing.T) {
	store := NewThreadStore()
	id1 := store.CreateThread(core.Thread{Title: "first", Body: "first"})
	id2 := store.CreateThread(core.Thread{Title: "second", Body: "second"})
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
	store := NewThreadStore()
	_, err := store.ReadThread(1)

	if err == nil {
		t.Fail()
	}
}

func TestCreateAndReadAllThreads(t *testing.T) {
	store := NewThreadStore()
	store.CreateThread(core.Thread{Title: "first"})
	store.CreateThread(core.Thread{Title: "second"})
	threads := store.ReadAllThreads()

	if len(threads) != 2 {
		t.Errorf("Expected 2 but got %d", len(threads))
	}
}
