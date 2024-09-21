package main

import (
	"fmt"
)

func main() {

	fmt.Println("hi")
	var a int
	println(a) //default value is 0
	a = 7
	println("new value of a is :", a)
	b := 5
	println("new value of b is :", b)
	c, d, e, f := "cat", 45.788, true, ""
	println("new value of c is :", c)
	println("new value of d is :", d)
	println("new value of e is :", e)
	println("new value of f is :", f)
	msg := fmt.Sprintf("this is string : %s \nthis is an integer : %d \nthis is an float : %f\nthis is an float till 2 decimal places by .2 : %.2f", "hello", 6, 5.7, d)
	println(msg)
	var x int = 77   //int is a general integer type used for arithmetic and counting.
	var y rune = 'k' //rune is an alias for int32 and is specifically used to represent Unicode characters.
	var z rune = 7786
	fmt.Printf("value of int as char %c and its type is %T \nvalue of rune as char %c and its type is %T \nvalue of rune as int %d \nvalue of rune as char %c", x, x, y, y, y, z)
	const l = "constant"
	// l = "change" this is not possible as this is an constant
	println(l)
	str := fmt.Sprintf("Hello, %s!", "World sprintf")
	fmt.Println(str)                           // Output: Hello, World!
	fmt.Printf("Hello, %s!\n", "World printf") // Output: Hello, World!
	// Printf: Formats and prints the output to the standard output (usually the console).
	//Sprintf: Formats the string but returns it instead of printing it. You can then use the string for further processing or store it in a variable.

}
