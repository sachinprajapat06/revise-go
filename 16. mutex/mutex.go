package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var k, counter int

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion
	for i := 1; i <= 3; i++ {
		k += id * 3
	}
	ch <- k // Send result to channel
}

func worker2(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion
	k = id * 2
	ch <- k // Send result to channel
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// Lock the mutex before accessing the shared counter
	mutex.Lock()
	counter++
	mutex.Unlock() // Unlock after modifying the counter
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 6) // Buffered channel
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
		wg.Add(1)
		go worker2(i, ch, &wg)
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	close(ch)
	for val := range ch {
		fmt.Println("1", val) // Prints values from 0 to 4
	}
	fmt.Println("Final Counter Value:", counter)
}

// ***********************************Explanation:
// Shared Data: The counter is shared between all goroutines.
// Mutex: The mutex ensures that only one goroutine can increment the counter at a time.
// Lock and Unlock: Before modifying the counter, each goroutine locks the mutex, and after modifying the counter, it unlocks it, allowing another waiting goroutine to lock the mutex and proceed.

// ***********************************Summary:
// Mutex provides a way to lock shared resources to prevent multiple goroutines from accessing them simultaneously.
// sync.Mutex is used for exclusive locking, and sync.RWMutex allows multiple readers but only one writer at a time.
// Use mutexes to avoid race conditions, but be careful to avoid deadlocks and unnecessary locking.

// **************************************Common Mistakes with Mutexes
// Forgetting to Unlock: Always ensure that the mutex is unlocked after locking, even in cases where the function may exit early due to an error. Use defer to ensure the mutex is unlocked properly:
// Deadlocks: Deadlocks can occur when two or more goroutines are stuck waiting for each other to release a lock. This can happen if one goroutine holds a lock and tries to acquire another lock that is held by another goroutine.
// Holding the Lock for Too Long: Minimize the amount of code executed between locking and unlocking a mutex. Holding a lock for too long can reduce concurrency and performance.
