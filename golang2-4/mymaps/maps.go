package mymaps

import "sync"

type SetMutex struct {
	sync.Mutex
	mm map[float64]struct{}
}

// NewSetMutex ...
func NewSetMutex() *SetMutex {
	return &SetMutex{
		mm: map[float64]struct{}{},
	}
}

func (s *SetMutex) Add(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *SetMutex) Has(i float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type SetRMutex struct {
	sync.RWMutex
	mm map[float64]struct{}
}

func NewSetRMutex() *SetRMutex {
	return &SetRMutex{
		mm: map[float64]struct{}{},
	}
}

func (s *SetRMutex) Add(i float64) {
	s.Lock()
	defer s.Unlock()
	s.mm[i] = struct{}{}

}

func (s *SetRMutex) Has(i float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}
