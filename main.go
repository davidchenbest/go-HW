package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func main() {
	c := Node[int]{val: 3}
	b := Node[int]{val: 2, next: &c}
	a := Node[int]{val: 1, next: &b}
	// d := Node[int]{val: 1}
	// e := Node[int]{val: 1}

	fmt.Println(a.next.next)
}
