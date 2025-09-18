package main

import (
	"errors"
	"fmt"
)

// Stack struct using fixed-size array
type Stack struct {
	data [100]int // fixed size array
	top  int      // index of the top element
}
// Initializes a new stack
func NewStack() *Stack {
	return &Stack{top: -1}
}
// Push adds an item to the top of the stack
func (s *Stack) Push(value int) error {
	if s.top >= len(s.data)-1 {
		return errors.New("stack overflow")
	}
	s.top++
	s.data[s.top] = value
	return nil
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, error) {
	if s.top < 0 {
		return 0, errors.New("stack underflow")
	}
	value := s.data[s.top]
	s.top--
	return value, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func main() {
	stack := NewStack()
	// Test pushing elements
	fmt.Println("Pushing values onto stack:")
	for i := 1; i <= 5; i++ {
		err := stack.Push(i * 10)
		if err != nil {
			fmt.Println("Push Error:", err)
		} else {
			fmt.Printf("Pushed: %d\n", i*10)
		}
	}
	// Test popping elements
	fmt.Println("\nPopping values from stack:")
	for i := 0; i < 5; i++ {
		value, err := stack.Pop()
		if err != nil {
			fmt.Println("Pop Error:", err)
		} else {
			fmt.Printf("Popped: %d\n", value)
		}
	}
}
