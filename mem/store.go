package mem

import (
	"errors"
	"strconv"

	"github.com/sandfort/goard/core"
)

type ThreadStore struct {
	Threads map[int]core.Thread
	Counter int
}

func NewThreadStore() *ThreadStore {
	return &ThreadStore{Threads: make(map[int]core.Thread), Counter: 0}
}

func (s *ThreadStore) CreateThread(p core.Thread) int {
	s.Counter += 1
	p.Id = s.Counter
	s.Threads[p.Id] = p
	return p.Id
}

func (s *ThreadStore) ReadThread(id int) (core.Thread, error) {
	p, ok := s.Threads[id]

	if !ok {
		return core.Thread{},
			errors.New("No thread with ID " + strconv.Itoa(id))
	}

	return p, nil
}

func (s *ThreadStore) ReadAllThreads() []core.Thread {
	var threads []core.Thread

	for _, thread := range s.Threads {
		threads = append(threads, thread)
	}

	return threads
}
