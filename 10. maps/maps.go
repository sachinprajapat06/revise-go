package main

import "fmt"

func main() {
	m := make(map[string]int)
	m2 := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	var m3 map[string]int
	m["Alice"] = 25 // Insert "Alice" with value 25
	age := m["Alice"]
	fmt.Println("value of alice is : ", age) // Outputs: 25

	value, exists := m3["Charlie"]
	if exists {
		fmt.Println("Key exists with value : ", value)
	} else {
		fmt.Println("Key does not exist")
	}
	delete(m2, "Alice") // For deleting an element

	nestedMap := make(map[string]map[string]int)
	nestedMap["outer"] = map[string]int{"inner": 42}
	fmt.Println(nestedMap["outer"]["inner"]) // Outputs: 42

	// var m map[string]int
	// m["key"] = 1 // Causes a runtime panic
	// Maps: provide fast key-based lookup and are unordered.
	// Slices: are ordered collections of elements that you access by index.
	// Maps are automatically garbage collected. When you delete a key, its memory is freed when no longer in use, but the overall capacity of the map remains unchanged unless you reallocate it.

}
