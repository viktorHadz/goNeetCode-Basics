package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	// Initialized when used with a type of my choosing. We got it finally! He-he
	stack := Stack[T]{}
	fmt.Printf("New empty stack creted\nStack: %+v\nValues: %+v", stack, stack.items)
	return &stack
}

// Receiver syntax example. This attaches the func to the struct

// Retrieves the current stack size
func (s *Stack[T]) GetStackSize() int {
	return len(s.items)
}

// Pushes a value of type T unto the stack
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Removes last elemet from stack and return false if stack empty
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastIdx := len(s.items) - 1
	v := s.items[lastIdx]

	s.items = s.items[:lastIdx]

	return v, true
}
