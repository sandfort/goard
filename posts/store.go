package posts

import (
	"strconv"
	"errors"
)

type MemoryStore struct {
	Posts map[int]Post
	Counter int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{Posts: make(map[int]Post), Counter: 0}
}

func (s *MemoryStore) CreatePost(p Post) int {
	s.Counter += 1
	p.Id = s.Counter
	s.Posts[p.Id] = p
	return p.Id
}

func (s *MemoryStore) ReadPost(id int) (Post, error) {
	p, ok := s.Posts[id]

	if !ok {
		return Post{}, errors.New("No Post with ID " + strconv.Itoa(id))
	}

	return p, nil
}

func (s *MemoryStore) ReadAllPosts() ([]Post) {
	var posts []Post

	for _, post := range s.Posts {
		posts = append(posts, post)
	}

	return posts
}
