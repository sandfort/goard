package mem

import (
	"errors"
	"github.com/sandfort/goard/core"
	"strconv"
)

type PostStore struct {
	Posts   map[int]core.Post
	Counter int
}

func NewPostStore() *PostStore {
	return &PostStore{Posts: make(map[int]core.Post), Counter: 0}
}

func (s *PostStore) CreatePost(p core.Post) int {
	s.Counter += 1
	p.Id = s.Counter
	s.Posts[p.Id] = p
	return p.Id
}

func (s *PostStore) ReadPost(id int) (core.Post, error) {
	p, ok := s.Posts[id]

	if !ok {
		return core.Post{},
			errors.New("No post with ID " + strconv.Itoa(id))
	}

	return p, nil
}

func (s *PostStore) ReadAllPosts() []core.Post {
	var posts []core.Post

	for _, post := range s.Posts {
		posts = append(posts, post)
	}

	return posts
}
