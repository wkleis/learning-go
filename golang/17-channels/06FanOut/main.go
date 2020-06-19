// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)
	for v := range c2 {
		fmt.Println("received: ", v)
	}

	fmt.Println("About to exit")

}

func populate(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
func fanOutIn(inChannel <-chan int, outChannel chan<- int) {
	const numberOfGoroutines = 10
	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)
	for i := 0; i < numberOfGoroutines; i++ {
		go func(id int) {
			for value := range inChannel {
				result := timeConsumingWork(value)
				fmt.Println("Goroutine", id, ": ", value, "->", result)
				outChannel <- result

			}
			wg.Done()
		}(i)
	}

	wg.Wait() //wait for all worker goroutines are complete
	close(outChannel)
}
func timeConsumingWork(input int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return input * rand.Intn(1000)
}
