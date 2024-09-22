package main

import "fmt"

func main() {
	// A slice is built on top of an array and has the following three components:
	// Pointer: Points to the starting position of the underlying array.
	// Length: The number of elements currently in the slice.
	// Capacity: The number of elements the slice can hold before it needs to allocate a new underlying array.

	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	slice2 := make([]int, 5)       // Creates a slice of length 5, capacity 5
	slice3 := make([]int, 3, 5)    // Creates a slice of length 3, capacity 5
	slice4 := []int{1, 2, 3, 4, 5} // Directly initialise

	fmt.Println("old slice is : ", slice) // slice is {20, 30, 40} created from an array
	fmt.Println("old length is : ", len(slice))
	fmt.Println("old capacity is : ", cap(slice)) // Output: 4 (capacity from index 1 to end of array)
	slice = append(slice, 1, 2, 3)                // No reallocation, slice capacity is enough
	fmt.Println("new 1 capacity is : ", cap(slice))
	slice = append(slice, 4, 6, 8) // Reallocation occurs here as capacity is exceeded
	fmt.Println("new length is : ", len(slice))
	fmt.Println("new 2 capacity is : ", cap(slice))
	fmt.Println("slice is : ", slice)
	fmt.Println("slice2 is : ", slice2)
	fmt.Printf("slice3 length is : %d \nslice3 capacity is : %d \n", len(slice3), cap(slice3)) //finding len, capacity
	fmt.Println("slice4 is : ", slice4)

	slice5 := arr[1:4]                             // {2, 3, 4}
	slice6 := arr[2:5]                             // {3, 4, 5}
	slice5[2] = 10                                 // slice1 becomes {2, 3, 10}, arr and slice2 also change
	fmt.Println("array is also changed : ", arr)   // Output: [1, 2, 3, 10, 5]
	fmt.Println("slice 6 also changed : ", slice6) // Output: {3, 10, 5}

	slice7 := []int{}
	fmt.Println("slice7 == nil : ", slice7 == nil) // Output: false
	var slice8 []int
	fmt.Println("slice8 == nil : ", slice8 == nil) // Output: true
}
