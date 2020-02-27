package syncstate

import (
	"covertest/state"
	"sync"
)

// State represents a thread-safe state object
type State struct {
	s  *state.State
	mu sync.RWMutex
}

// New creates a new thread-safe state object
func New(s *state.State) *State {
	return &State{s: s}
}

// Get returns the state value
func (s *State) Get() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.s.Get()
}

// SetInt sets the state value from an integer
func (s *State) SetInt(i int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.s.SetInt(i)
}
