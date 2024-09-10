package main

import "fmt"

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
}
