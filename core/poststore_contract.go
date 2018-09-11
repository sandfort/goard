package core

import "testing"

func NewPostStoreContract(store PostStore) *fixture {
	return &fixture{store: store}
}

type fixture struct {
	store PostStore
}

type PostStoreContract interface {
	CreateAndReadPostContract(t *testing.T)
	CreateAndReadMultiplePostsContract(t *testing.T)
	ReadPostReturnsErrorWhenIdDoesNotExistContract(t *testing.T)
	ReadByThreadIdContract(t *testing.T)
}

func (f *fixture) CreateAndReadPostContract(t *testing.T) {
	body := "The Body"
	id := f.store.CreatePost(Post{Body: body})
	result, _ := f.store.ReadPost(id)
	if result.Body != body {
		t.Errorf("Was expecting body to be %q but was %q", body, result.Body)
	}
}

func (f *fixture) CreateAndReadMultiplePostsContract(t *testing.T) {
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

func (f *fixture) ReadPostReturnsErrorWhenIdDoesNotExistContract(t *testing.T) {
	_, err := f.store.ReadPost(1)

	if err == nil {
		t.Error("Expected error")
	}
}

func (f *fixture) ReadByThreadIdContract(t *testing.T) {
	id := f.store.CreatePost(Post{ThreadId: 1})
	f.store.CreatePost(Post{ThreadId: 2})
	f.store.CreatePost(Post{ThreadId: 2})

	t1Posts := f.store.ReadByThreadId(1)

	if len(t1Posts) != 1 {
		t.Errorf("Expected thread to have 1 post but found %d", len(t1Posts))
	}

	if t1Posts[0].Id != id {
		t.Errorf("Expected to find post with ID %d but found ID %d", id, t1Posts[0].Id)
	}
}
