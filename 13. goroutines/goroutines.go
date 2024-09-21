package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	fmt.Println("started")
	a := make(chan string, 10)

	fmt.Println("default value of a channel :", a)
	go twice("hi", a, wg)

	result := <-a // getting values of channels into var
	r2 := <-a
	r3 := <-a
	r4 := <-a
	fmt.Println("here is the result :", result)
	fmt.Println("here is the result1 :", r2)
	fmt.Println("here is the result2 :", r3)
	fmt.Println("here is the result3 :", r4)

	for res := range a { // directly print them
		println(res)
	}
	wg.Wait()

}
func twice(str string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	str = str + str
	wg.Add(1)
	go four(str, ch, wg)
	six(str, ch)
	fmt.Println("result  in goroutine :", str)
}

func four(str string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	str = str + str
	ch <- str
	fmt.Println("result  in goroutine :", str)
	close(ch)
}

func six(str string, ch chan string) {
	str = str + str + str
	ch <- str
	ch <- "a"
	ch <- "b"
	fmt.Println("result in  :", str)
}
