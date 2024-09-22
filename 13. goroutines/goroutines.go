package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1) // Increment the counter for the goroutine

	fmt.Println("started")
	a := make(chan string, 10)
	// defer close(a) //we cant do this here as this closes after goroutines

	fmt.Println("default value of a channel :", a)
	go twice("hi", a, wg)

	result := <-a // getting values of channels into var
	r2 := <-a
	r3 := <-a
	r4 := <-a
	r5 := <-a
	fmt.Println("here is the result :", result)
	fmt.Println("here is the result1 :", r2)
	fmt.Println("here is the result2 :", r3)
	fmt.Println("here is the result3 :", r4)
	fmt.Println("here is the result4 :", r5)

	for res := range a { // directly print them
		println("responses :", res)
	}
	wg.Wait() // Block the program until the counter is decremented to 0

}
func twice(str string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	//defer close(ch) // we can not close it here as new goroutine is using it and our funtion may end before that goroutine
	str = str + str // string get doubled
	ch <- str       // channel will get result

	wg.Add(1)            // Adding another waitgroup as we are calling an other gorotine here
	go four(str, ch, wg) // calling another goroutine
	six(str, ch)
	fmt.Println("result  in goroutine twice :", str)
}

func four(str string, ch chan string, wg *sync.WaitGroup) {
	defer close(ch) // channel will get close at last
	defer println("chanel is closed")
	defer wg.Done() // Decrement the counter when the goroutine completes
	//defer close(ch) // channel will get close at last
	str = str + str
	ch <- str // channel will get result
	fmt.Println("result  in goroutine four :", str)
	//close(ch) // this is the last channel value sending so we close this
	// here we close this because not closing a channel will give asleap goroutine
}

func six(str string, ch chan string) {
	str = str + str + str
	ch <- str
	ch <- "a"
	ch <- "b"
	fmt.Println("result in  :", str)
	// close(ch)
	// here we cant close channel as it is executed before then the other goroutine
}

// runtime.GOMAXPROCS(4) // Allow up to 4 OS threads
// Concurrency is about dealing with many things at once, which means multiple tasks are in progress simultaneously, though not necessarily at the same time.
// Parallelism is when multiple tasks are running at the exact same time, utilizing multiple processors/cores.
// Channels: These are pipes that allow goroutines to communicate with each other by sending and receiving values. Channels are crucial for safe and easy communication between goroutines.

// ********************* Best Practices with Goroutines
// Avoid long-running goroutines: Only use goroutines when necessary, and avoid letting them run indefinitely unless there is a clear reason.
// Synchronize using channels and other synchronization primitives: Always ensure goroutines communicate and synchronize using appropriate mechanisms like channels, mutexes, or waitgroups.
// Avoid race conditions: When goroutines access shared data, use mutexes or channel-based communication to avoid race conditions.

// ********************* Practical Use Cases
// Goroutines are useful in a wide variety of applications:
// Web servers: Handling multiple requests concurrently.
// Task scheduling: Running background jobs concurrently with other operations.
// Parallel processing: Performing tasks on multiple cores to improve performance.
