package main

// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod
// 2nd demo is copied from Rob Pike https://talks.golang.org/2012/concurrency.slide#27
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fanInDemo2()

}

func fanInDemo() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	fanInChannel := make(chan int)

	go send(evenChannel, oddChannel)
	go receive(evenChannel, oddChannel, fanInChannel)

	for v := range fanInChannel {
		fmt.Println(v)
	}

	fmt.Println("About to exit ...")
}

func send(evenChannel, oddChannel chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}
	close(evenChannel)
	close(oddChannel)
}

func receive(evenChannel, oddChannel chan int, fanInChannel chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range evenChannel {
			fanInChannel <- v
		}
		wg.Done()
	}()
	go func() {

		for v := range oddChannel {
			fanInChannel <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanInChannel)

}

// the following code is the concurrency example from Rob Pike
// https://talks.golang.org/2012/concurrency.slide#27
func fanInDemo2() {

	jerry := sourceChannel("Jerry")
	tom := sourceChannel("Tom")

	channel := fanIn(jerry, tom)
	for i := 0; i < 20; i++ {
		fmt.Println(<-channel)
	}
	fmt.Println("OK, leaving ...")
}

func sourceChannel(name string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("%s %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return channel
}

func fanIn(input1 <-chan string, input2 <-chan string) <-chan string {
	outChannel := make(chan string)
	go func() {
		for {
			outChannel <- <-input1
		}

	}()
	go func() {
		for {
			outChannel <- <-input2
		}
	}()
	return outChannel
}
