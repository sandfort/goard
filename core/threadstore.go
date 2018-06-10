package core

import (
	"errors"
	"strconv"
)

type ThreadMemoryStore struct {
	Threads map[int]Thread
	Counter int
}

func NewThreadMemoryStore() *ThreadMemoryStore {
	return &ThreadMemoryStore{Threads: make(map[int]Thread), Counter: 0}
}

func (s *ThreadMemoryStore) CreateThread(p Thread) int {
	s.Counter += 1
	p.Id = s.Counter
	s.Threads[p.Id] = p
	return p.Id
}

func (s *ThreadMemoryStore) ReadThread(id int) (Thread, error) {
	p, ok := s.Threads[id]

	if !ok {
		return Thread{},
			errors.New("No thread with ID " + strconv.Itoa(id))
	}

	return p, nil
}

func (s *ThreadMemoryStore) ReadAllThreads() []Thread {
	var threads []Thread

	for _, thread := range s.Threads {
		threads = append(threads, thread)
	}

	return threads
}
