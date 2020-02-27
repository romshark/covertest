package state

import (
	"fmt"
)

// State represents a stateful object
type State struct {
	state string
}

// State types
const (
	Unknown = "unknown"
	Small   = "small"
	Average = "average"
	Great   = "great"
)

// Get returns the state value
func (s *State) Get() string {
	return s.state
}

// SetInt sets the state value from an integer
func (s *State) SetInt(i int) error {
	switch {
	case i < 10:
		s.state = Small
	case i == 15:
		s.state = Average
	case i > 20:
		s.state = Great
	default:
		return fmt.Errorf("unknown state id: %d", i)
	}
	return nil
}
