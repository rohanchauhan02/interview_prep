package stacks

import "fmt"

// 6. Implement a stack supporting push, pop, and min operations.

type stack struct {
	data     []int
	minStack []int
	min      int
	size     int
}

func NewStack() stack {
	return stack{
		data:     []int{},
		minStack: []int{},
		size:     0,
	}
}

func (s *stack) Push(val int) {
	if s.IsEmpty() {
		s.data = append(s.data, val)
		s.min = val
		s.minStack = append(s.minStack, val)
		s.size++
		return
	}
	if s.min < val {
		s.min = val
	}
	s.data = append(s.data, val)
	s.minStack = append(s.minStack, s.min)
	s.size++
}

func (s *stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}
	s.data = s.data[:len(s.data)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
	s.min = s.topMin()
	s.size--
}

func (s *stack) Top() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *stack) topMin() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.minStack[len(s.minStack)-1]
}

func (s *stack) Min() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.min
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}
