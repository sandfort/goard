package mem

import (
	"github.com/sandfort/goard/core"
	"testing"
)

func TestCreateAndReadPost(t *testing.T) {
	store := NewPostStore()
	id := store.CreatePost(core.Post{
		Title: "yo",
		Body:  "the body of the post",
	})
	result, _ := store.ReadPost(id)

	if result.Title != "yo" || result.Body != "the body of the post" {
		t.Fail()
	}
}

func TestCreateAndReadMultiplePosts(t *testing.T) {
	store := NewPostStore()
	id1 := store.CreatePost(core.Post{Title: "first", Body: "first"})
	id2 := store.CreatePost(core.Post{Title: "second", Body: "second"})
	p1, _ := store.ReadPost(id1)
	p2, _ := store.ReadPost(id2)

	if p1.Title != "first" {
		t.Fail()
	}

	if p2.Title != "second" {
		t.Fail()
	}
}

func TestReadReturnsErrorWhenIdDoesNotExist(t *testing.T) {
	store := NewPostStore()
	_, err := store.ReadPost(1)

	if err == nil {
		t.Fail()
	}
}

func TestCreateAndReadAllPosts(t *testing.T) {
	store := NewPostStore()
	store.CreatePost(core.Post{Title: "first"})
	store.CreatePost(core.Post{Title: "second"})
	posts := store.ReadAllPosts()

	if len(posts) != 2 {
		t.Errorf("Expected 2 but got %d", len(posts))
	}
}
