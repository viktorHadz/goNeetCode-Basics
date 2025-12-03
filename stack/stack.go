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

func (s *Stack[T]) Pop() (T, bool) {
	last := len(s.items) - 1
	newSlice := s.items[0:last]
	lastEl := s.items[last:]

	if lastEl != nil {
		s.items = newSlice
		return s.items[last], true
	} else {
		return s.items[last], false
	}
}
