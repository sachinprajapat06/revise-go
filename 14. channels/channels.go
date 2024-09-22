package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion
	result := id * 2
	ch <- result // Send result to channel
}

func worker2(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion
	result := id * 2
	ch <- result // Send result to channel
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)  // Buffered channel
	ch2 := make(chan int, 3) // UnBuffered channel
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
		wg.Add(1)
		go worker2(i, ch2, &wg)
	}

	wg.Wait()
	close(ch)
	close(ch2)
	for val := range ch2 {
		fmt.Println("2", val) // Prints values from 0 to 4
	}
	for val := range ch {
		fmt.Println("1", val) // Prints values from 0 to 4
	}

}
