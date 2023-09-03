package main

import "fmt"

func main() {
	// Defer a function that recovers from panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate a panic
	divideByZero()

	// This will not be reached due to the panic and recovery
	fmt.Println("End of the program")
}

func divideByZero() {
	panic("not allowed")
}
