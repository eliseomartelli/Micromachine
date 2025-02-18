package micromachine

import (
	"fmt"
	"sync"
)

// Micromachine is a generic state machine that supports any comparable type
// for states.
type Micromachine[T comparable] struct {
	mu          sync.Mutex
	state       T
	transitions map[T]map[T]func() error
}

// NewMicromachine creates a new state machine with the given initial state.
func NewMicromachine[T comparable](initialState T) *Micromachine[T] {
	return &Micromachine[T]{
		state:       initialState,
		transitions: make(map[T]map[T]func() error),
	}
}

// AddTransition adds a transition from one state to another with an optional
// action. The action is a function that will be executed during the
// transition.
func (sm *Micromachine[T]) AddTransition(from, to T, action func() error) *Micromachine[T] {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sm.transitions[from] == nil {
		sm.transitions[from] = make(map[T]func() error)
	}
	sm.transitions[from][to] = action
	return sm
}

// CanTransition checks if a transition to the given state is valid from the
// current state.
func (sm *Micromachine[T]) CanTransition(to T) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	_, exists := sm.transitions[sm.state][to]
	return exists
}

// Transition attempts to transition to the given state.
// If the transition is valid and an action is defined, the action is executed.
// Returns an error if the transition is invalid.
func (sm *Micromachine[T]) Transition(to T) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if action, exists := sm.transitions[sm.state][to]; exists {
		if action != nil {
			if err := action(); err != nil {
				return err
			}
		}
		sm.state = to
		return nil
	}
	return fmt.Errorf(
		"invalid state transition from '%v' to '%v'",
		sm.state,
		to,
	)
}

// State returns the current state of the state machine.
func (sm *Micromachine[T]) State() T {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.state
}

// ValidTransitions returns a slice of valid states that can be transitioned
// to from the current state.
func (sm *Micromachine[T]) ValidTransitions() []T {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	valid := make([]T, 0, len(sm.transitions[sm.state]))
	for to := range sm.transitions[sm.state] {
		valid = append(valid, to)
	}
	return valid
}
