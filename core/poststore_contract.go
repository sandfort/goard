package core

import "testing"

func NewPostStoreContract(store PostStore) *fixture {
	return &fixture{store: store}
}

type fixture struct {
	store PostStore
}

type PostStoreContract interface {
	Verify(t *testing.T)
	createAndReadPostContract(t *testing.T)
	createAndReadMultiplePostsContract(t *testing.T)
	readPostReturnsErrorWhenIdDoesNotExistContract(t *testing.T)
	readByThreadIdContract(t *testing.T)
}

func (f *fixture) Verify(t *testing.T) {
	f.createAndReadPostContract(t)
	f.createAndReadMultiplePostsContract(t)
	f.readPostReturnsErrorWhenIdDoesNotExistContract(t)
	f.readByThreadIdContract(t)
}

func (f *fixture) createAndReadPostContract(t *testing.T) {
	body := "The Body"
	id := f.store.CreatePost(Post{Body: body})
	result, _ := f.store.ReadPost(id)
	if result.Body != body {
		t.Errorf("Was expecting body to be %q but was %q", body, result.Body)
	}
}

func (f *fixture) createAndReadMultiplePostsContract(t *testing.T) {
	body1 := "first"
	body2 := "second"

	id1 := f.store.CreatePost(Post{Body: body1})
	id2 := f.store.CreatePost(Post{Body: body2})

	post1, _ := f.store.ReadPost(id1)
	post2, _ := f.store.ReadPost(id2)

	if post1.Body != body1 {
		t.Errorf("Expected body to be %q but was %q", body1, post1.Body)
	}

	if post2.Body != body2 {
		t.Errorf("Expected body to be %q but was %q", body2, post2.Body)
	}
}

func (f *fixture) readPostReturnsErrorWhenIdDoesNotExistContract(t *testing.T) {
	_, err := f.store.ReadPost(99)

	if err == nil {
		t.Error("Expected error")
	}
}

func (f *fixture) readByThreadIdContract(t *testing.T) {
	id := f.store.CreatePost(Post{ThreadID: 1})
	f.store.CreatePost(Post{ThreadID: 2})
	f.store.CreatePost(Post{ThreadID: 2})

	t1Posts := f.store.ReadByThreadId(1)

	if len(t1Posts) != 1 {
		t.Errorf("Expected thread to have 1 post but found %d", len(t1Posts))
	}

	if t1Posts[0].ID != id {
		t.Errorf("Expected to find post with ID %d but found ID %d", id, t1Posts[0].ID)
	}
}
