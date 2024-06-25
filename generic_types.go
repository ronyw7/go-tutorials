package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Push adds a new element to the front of the list
func (l *List[T]) Push(value T) *List[T] {
	return &List[T]{next: l, val: value}
}

// Pop returns the first element from the list and returns its value
func (l *List[T]) Pop() (T, bool) {
	if l.next == nil {
		return l.val, false
	}
	return l.val, true
}

// PrintAll prints all elements in the list
func (l *List[T]) PrintAll() {
	for node := l; node != nil; node = node.next {
		fmt.Print(node.val, " ")
	}
	fmt.Println()
}

func main() {
	var head *List[int] = &List[int]{}
	head = head.Push(1)
	head = head.Push(2)
	head = head.Push(3)

	head.PrintAll()

	if value, ok := head.Pop(); ok {
		fmt.Printf("Popped: %v\n", value)
	}
}
