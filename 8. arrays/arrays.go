package main

import "fmt"

func main() {
	var arr [5]int                // array of 5 integers, default values are 0
	arr2 := [5]int{1, 2, 3, 4, 5} // an array initialized with specific values
	arr3 := [...]int{1, 2, 3}     // array size is inferred to be 3
	arr4 := [5]int{1, 2}          // array: [1, 2, 0, 0, 0]
	var defaultArr [3]int
	arr[1] = 50                                                                        // Modifies the second element to 50
	fmt.Println("arr 2st element is : ", arr[1])                                       // Output: 50
	fmt.Println("default values of an array ", defaultArr)                             // Output: [0, 0, 0]
	fmt.Println("arr2 1st element is : ", arr2[0])                                     // Output: 10
	fmt.Println("arr3 length is : ", len(arr3))                                        // Output: 5
	fmt.Printf("arr4 length is : %d \narr4 capacity is : %d \n", len(arr4), cap(arr4)) //finding len, capacity
	fmt.Println("arr == arr4 : ", arr == arr4)                                         // will be false because both have different values                                                      // Comparing 2 arrays

	//Fixed size: Arrays cannot grow or shrink.
	//Dynamic size: Slices, created from arrays, can dynamically resize, making them more flexible.
}
