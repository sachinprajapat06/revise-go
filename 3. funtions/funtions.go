package main

type rect struct {
	width  int
	height int
}

func main() {
	println("concat nested function output is ", concat(concat("s1", "s2"), "s3"))
	println("get val function will get us ")
	println(getVal())

	r := rect{}
	r.width = 5
	r.height = 4
	println("area of rectangle is ", r.area())

}

func concat(s1, s2 string) string {
	return s1 + s2
}

func getVal() (x, y int) {
	return // no need to return values
	// we can do this     return x, y
	// in this we are returning our choice variables not x,y    return 5, 6
}

func (r rect) area() int {
	return r.height * r.width
}
