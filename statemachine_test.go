package micromachine

import (
	"errors"
	"testing"
)

type ProcessState string

const (
	Idle    ProcessState = "idle"
	Running ProcessState = "running"
	Stopped ProcessState = "stopped"
)

func TestStateMachine_Transitions(t *testing.T) {
	sm := NewMicromachine(Idle)

	sm.
		AddTransition(Idle, Running, nil).
		AddTransition(Running, Stopped, nil).
		AddTransition(Stopped, Idle, nil)

	if sm.State() != Idle {
		t.Errorf("expected initial state to be %v, got %v", Idle, sm.State())
	}

	if err := sm.Transition(Running); err != nil {
		t.Errorf("unexpected error during transition: %v", err)
	}
	if sm.State() != Running {
		t.Errorf("expected state to be %v, got %v", Running, sm.State())
	}

	if err := sm.Transition(Stopped); err != nil {
		t.Errorf("unexpected error during transition: %v", err)
	}
	if sm.State() != Stopped {
		t.Errorf("expected state to be %v, got %v", Stopped, sm.State())
	}

	if err := sm.Transition(Idle); err != nil {
		t.Errorf("unexpected error during transition: %v", err)
	}
	if sm.State() != Idle {
		t.Errorf("expected state to be %v, got %v", Idle, sm.State())
	}
}

func TestStateMachine_InvalidTransition(t *testing.T) {
	sm := NewMicromachine(Idle)

	sm.AddTransition(Idle, Running, nil)

	err := sm.Transition(Stopped)
	if err == nil {
		t.Errorf("expected error for invalid transition, got nil")
	}
}

func TestStateMachine_CanTransition(t *testing.T) {
	sm := NewMicromachine(Idle)

	sm.AddTransition(Idle, Running, nil)
	sm.AddTransition(Running, Stopped, nil)

	if !sm.CanTransition(Running) {
		t.Errorf("expected CanTransition(Running) to return true")
	}
	if sm.CanTransition(Stopped) {
		t.Errorf("expected CanTransition(Stopped) to return false from Idle state")
	}
}

func TestStateMachine_TransitionWithAction(t *testing.T) {
	sm := NewMicromachine(Idle)

	actionCalled := false
	sm.AddTransition(Idle, Running, func() error {
		actionCalled = true
		return nil
	})

	err := sm.Transition(Running)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !actionCalled {
		t.Errorf("expected action to be called, but it wasn't")
	}
}

func TestStateMachine_TransitionWithError(t *testing.T) {
	sm := NewMicromachine(Idle)

	sm.AddTransition(Idle, Running, func() error {
		return errors.New("failed action")
	})

	err := sm.Transition(Running)
	if err == nil || err.Error() != "failed action" {
		t.Errorf("expected 'failed action' error, got %v", err)
	}

	if sm.State() != Idle {
		t.Errorf("expected state to remain %v, but got %v", Idle, sm.State())
	}
}
