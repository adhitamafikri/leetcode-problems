package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (l *List[T]) Push(val T) {
	fmt.Println("Pushing val: ", val)

	if l.tail == nil {
		fmt.Println("l.tail == nil", val)
		l.head = &element[T]{next: nil, val: val}
		l.tail = l.head
	} else {
		fmt.Println("l.tail != nil", val)
		l.tail.next = &element[T]{next: nil, val: val}
		l.tail = l.tail.next
	}
}

func (l *List[T]) AllValues() []T {
	var values []T

	for e := l.head; e != nil; e = e.next {
		values = append(values, e.val)
	}

	return values
}

func main() {
	sVal := []string{"Adhitama", "Fikri", "K"}
	var sList List[string]

	for _, val := range sVal {
		sList.Push(val)
	}

	fmt.Println("This is the list", sList)
	fmt.Println("All values from the list: ", sList.AllValues())
}
