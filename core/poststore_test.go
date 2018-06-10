package core

import "testing"

func TestCreateAndReadPost(t *testing.T) {
	store := NewPostMemoryStore()
	body := "The Body"
	id := store.CreatePost(Post{Body: body})
	result, _ := store.ReadPost(id)
	if result.Body != body {
		t.Errorf("Was expecting body to be %q but was %q", body, result.Body)
	}
}

func TestCreateAndReadMultiplePosts(t *testing.T) {
	store := NewPostMemoryStore()

	body1 := "first"
	body2 := "second"

	id1 := store.CreatePost(Post{Body: body1})
	id2 := store.CreatePost(Post{Body: body2})

	post1, _ := store.ReadPost(id1)
	post2, _ := store.ReadPost(id2)

	if post1.Body != body1 {
		t.Errorf("Expected body to be %q but was %q", body1, post1.Body)
	}

	if post2.Body != body2 {
		t.Errorf("Expected body to be %q but was %q", body2, post2.Body)
	}
}

func TestReadPostReturnsErrorWhenIdDoesNotExist(t *testing.T) {
	store := NewPostMemoryStore()
	_, err := store.ReadPost(1)

	if err == nil {
		t.Error("Expected error")
	}
}

func TestReadByThreadId(t *testing.T) {
	postStore := NewPostMemoryStore()

	id := postStore.CreatePost(Post{ThreadId: 1})
	postStore.CreatePost(Post{ThreadId: 2})
	postStore.CreatePost(Post{ThreadId: 2})

	t1Posts := postStore.ReadByThreadId(1)

	if len(t1Posts) != 1 {
		t.Errorf("Expected thread to have 1 post but found %d", len(t1Posts))
	}

	if t1Posts[0].Id != id {
		t.Errorf("Expected to find post with ID %d but found ID %d", id, t1Posts[0].Id)
	}
}
