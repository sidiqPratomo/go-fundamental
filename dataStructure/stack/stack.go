package main

import "fmt"

type Stack struct {
	items []int
}

func (s Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item
}

func (s Stack) Peek() int {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}

	return s.items[len(s.items)-1]
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Is stack empty?", stack.IsEmpty())

	fmt.Println("Top item:", stack.Peek())

	fmt.Println("Popped item:", stack.Pop())
	fmt.Println("Popped item:", stack.Pop())
	fmt.Println("Popped item:", stack.Pop())

	fmt.Println("Is stack empty?", stack.IsEmpty())

	// Trying to pop from an empty stack will cause a panic
	// fmt.Println("Popped item:", stack.Pop())
}
