package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Numeric is a constraint that allows any numeric type
type Numeric interface {
	constraints.Integer | constraints.Float
}

// Generic function that works for any type that supports comparison
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(3, 4))            // Works for integers
	fmt.Println(Max(3.14, 2.71))      // Works for floats
	fmt.Println(Max("apple", "pear")) // Works for strings
}

// *******************************Explanation:
// constraints.Ordered is a type constraint that includes all types that can be compared using the <, >, <=, and >= operators. This includes int, float, and string.

// ******************************Benefits of Generics:
// Code Reusability: You can write a single function or type that works with any data type, reducing code duplication.
// Type Safety: Generic code is type-safe, meaning type checks happen at compile time, preventing runtime type errors.
// Performance: Go's generics are optimized to maintain performance, unlike some languages where generic types can lead to performance penalties.

// ******************************Key Points to Remember:
// Generics use type parameters ([T any]).
// Type constraints are defined using interfaces.
// You can use generics with functions, methods, structs, and more.
// Type inference allows you to omit type parameters in many cases.
// Generic code is compiled to work efficiently without runtime overhead.

// In summary, generics in Go enable developers to write more flexible and reusable code while preserving the simplicity and performance benefits Go is known for.
