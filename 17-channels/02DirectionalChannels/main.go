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
	fmt.Printf("Sender: %T\n", channel)
	channel <- 42
}
func receiver(channel <-chan int) {
	fmt.Printf("Receiver: %T\n", channel)
	fmt.Println("Receiver:", <-channel)
}
