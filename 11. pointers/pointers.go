package main

import (
	"fmt"
	"sync"
)

// &: the address of operator, returns the memory address of a variable.
// *: the dereference operator, retrieves the value from the memory address stored in the pointer.

type Person struct {
	name string
	age  int
}

func changeName(p *Person) {
	p.name = "Alice"
}

func increment(mu sync.Mutex, counter *int) {
	*counter += 5
	println("counter value : ", counter)
	println("&counter value : ", &counter)
	println("*counter value : ", *counter)
	mu.Lock()
	*counter += 5
	co := *counter
	println("value of co ", co) // Modifying shared counter
	mu.Unlock()
}

func main() {
	x := 42      // Allocate memory for int on stack
	ptr := &x    // ptr holds the address of x
	pptr := &ptr // pptr is a pointer to ptr

	var intPtr *int                         // intPtr is a pointer to an int
	fmt.Println("value of ptr : ", *ptr)    // Dereferencing pptr twice gives the value of x (5)
	fmt.Println("value of pptr : ", **pptr) // Dereferencing pptr twice gives the value of x (5)

	person := Person{"Bob", 30}
	changeName(&person)      // Passing the pointer to the struct
	fmt.Println(person.name) // Outputs: Alice

	// Nil Pointers and Pointer Safety
	var p *int
	if p != nil {
		fmt.Println(*p) // Safe to dereference
	}

	println("done")
	//Using mutex
	var mu sync.Mutex
	intPtr = &x
	increment(mu, intPtr)

	// Arrays: Fixed-size, passed by value. Pointers to arrays allow you to manipulate the array in place.
	// Slices: Unlike arrays, slices already behave like a reference to an underlying array, so passing a slice by value does not create a copy of the entire data.
}
