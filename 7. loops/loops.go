package main

import "fmt"

func main() {

	// simple for loop
	for i := 0; i < 7; i++ {
		if i == 3 {
			continue // Skip iteration when i is 3
		}
		if i == 5 {
			break // Exit the loop when i is 5
		}
		fmt.Println(i)
	}

	// do while loop
	i := 0
	for i < 5 { // equivalent to while (i < 5) in other languages
		fmt.Println(i)
		i++
	}

	// This is an infinite loop
	for {
		// This is an infinite loop
	}

	// range loops like python
	nums := []int{10, 20, 30, 40}
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// string loop
	str := "Hello"
	for i, v := range str {
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	// map loop
	capitals := map[string]string{"France": "Paris", "Japan": "Tokyo"}
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}

	//outerloop an advance loop topic
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OuterLoop // Break out of the outer loop
			}
			fmt.Println(i, j)
		}
	}

}
