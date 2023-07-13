package main

import "fmt"

type Node2 struct {
	next *Node2
	val  any
}

type Node[T any] struct {
	next *Node[T]
	val  T
}

func main() {
	c := Node[int]{val: 3}
	b := Node[int]{val: 2, next: &c}
	a := Node[int]{val: 1, next: &b}

	g := Node2{val: 6}
	f := Node2{val: "5", next: &g}
	e := Node2{val: 4, next: &f}

	fmt.Println(a.next)
	fmt.Println(e.next)
}
