package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "Beep boop, I am " + r.Model
}

func greet(s Speaker) {
	fmt.Println(s.Speak())
}

func describe(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Integer:", v)
	default:
		fmt.Println("Unknown type")
	}
}

//Interfaces in Go define behavior by specifying method signatures that types must implement.
//Go interfaces are implicit, meaning types automatically implement an interface if they provide the required methods.
//The empty interface (interface{}) can hold any type.
//You can use type assertions and type switches to extract or inspect the underlying type of an interface at runtime.
//Interfaces are a powerful tool for polymorphism, decoupling, and writing flexible, reusable code in Go.

func main() {
	p := Person{Name: "Alice"}
	r := Robot{Model: "RX-78"}

	greet(p) // Output: Hello, my name is Alice
	greet(r) // Output: Beep boop, I am RX-78

	describe("Hello") // Output: String: Hello
	describe(42)      // Output: Integer: 42
	describe(3.14)    // Output: Unknown type
}
