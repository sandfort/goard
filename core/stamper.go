package core

type Stamper interface {
	Stamp() int
}

func NewIncrementingStamper() Stamper {
	return &IncrementingStamper{Count: 0}
}

type IncrementingStamper struct {
	Count int
}

// Stamp returns 1 on the first call after construction, and on subsequent calls, the previous stamp plus 1.
func (s *IncrementingStamper) Stamp() int {
	s.Count++
	return s.Count
}
