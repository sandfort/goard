package core

import (
	"fmt"
)

type PostMemoryStore struct {
	Posts   map[int]Post
	Counter int
}

func NewPostMemoryStore() PostStore {
	return &PostMemoryStore{Posts: make(map[int]Post), Counter: 0}
}

func (s *PostMemoryStore) CreatePost(post Post) int {
	s.Counter += 1
	post.Id = s.Counter
	s.Posts[post.Id] = post
	return post.Id
}

func (s *PostMemoryStore) ReadPost(id int) (Post, error) {
	p, ok := s.Posts[id]

	if !ok {
		return Post{}, fmt.Errorf("No post with id %d", id)
	}

	return p, nil
}

func (s *PostMemoryStore) ReadByThreadId(tid int) []Post {
	var ps []Post

	for _, p := range s.Posts {
		if p.ThreadId == tid {
			ps = append(ps, p)
		}
	}

	return ps
}
