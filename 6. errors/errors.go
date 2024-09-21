package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
		// here we can return it but cant format it

		//return 0, fmt.Errorf("cannot divide %f by zero", a)
		//using errorf we can format and return an error
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0) // in go we have to explacitly check errors
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// error type: Go’s built-in interface for handling errors, with an Error() method that returns a string.
// Explicit error handling: Functions that can fail return an error, which the caller must check.
// Creating errors: Use errors.New, fmt.Errorf, or define custom error types.
// Error vs Panic: Errors are used for recoverable conditions, while panics are used for critical errors.
// panic("something went wrong")
