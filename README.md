# Micromachine

This package provides a generic state machine implementation in Go. It allows
you to define states and transitions between them, with optional actions that
can be executed during transitions. The state machine is thread-safe, ensuring
that concurrent access is handled correctly.


## Features

- Supports any comparable type for states.
- Thread-Safe
- Define custom actions to be executed during state transitions.
- Checks if a transition is valid before attempting it.

## Usage

### Creating a State Machine

```go
package main

import (
	"fmt"
	"github.com/eliseomartelli/micromachine"
)

type ProcessState string

const (
	Idle    ProcessState = "idle"
	Running ProcessState = "running"
	Stopped ProcessState = "stopped"
)

func main() {
	sm := micromachine.NewMicromachine(Idle)

	sm.
		AddTransition(Idle, Running, nil).
		AddTransition(Running, Stopped, nil).
		AddTransition(Stopped, Idle, nil)

	fmt.Println("Initial State:", sm.State())

	if err := sm.Transition(Running); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current State:", sm.State())

	if err := sm.Transition(Stopped); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current State:", sm.State())

	if err := sm.Transition(Idle); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current State:", sm.State())
}
```

### Adding Transitions with Actions

```go
sm.AddTransition(Idle, Running, func() error {
    fmt.Printf("Hello, world!")
	return nil
})

err := sm.Transition(Running)
if err != nil {
	fmt.Println("Error:", err)
}
```

### Checking Valid Transitions

```go
if sm.CanTransition(Running) {
	fmt.Println("Transition to Running is valid")
} else {
	fmt.Println("Transition to Running is not valid")
}
```

### Handling Invalid Transitions

```go
err := sm.Transition(Stopped)
if err != nil {
	fmt.Println("Error:", err)
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE)
file for details.

