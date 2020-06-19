package main
// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	fmt.Printf("%T\n", channel)
	go sender(channel)
	receiver(channel)

}

func sender(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}
func receiver(channel <-chan int) {
	for value := range channel {
		fmt.Println("Receiver:", value)
	}

}
