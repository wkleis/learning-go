package main
// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import "fmt"

func main() {
	channelsBlockDemo()
	channelsBlockDemoBuffered()
}

func channelsBlockDemo() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

}
func channelsBlockDemoBuffered() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)

}
