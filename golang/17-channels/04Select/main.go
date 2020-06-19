package main

// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import "fmt"

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan int)

	go send(evenChannel, oddChannel, quitChannel)
	receive(evenChannel, oddChannel, quitChannel)

	fmt.Println("About to exit ...")

}

func send(evenChannel, oddChannel, quitChannel chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}
	close(evenChannel)
	close(oddChannel)
	quitChannel <- 1
	close(quitChannel)
}

func receive(evenChannel, oddChannel, quitChannel <-chan int) {
	for {
		select {
		case value, ok := <-evenChannel:
			if ok {
				fmt.Println("from even channel:", value)
			} else {
				fmt.Println("Not OK value from even channel (closed)")
			}
		case value, ok := <-oddChannel:
			if ok {
				fmt.Println("from odd channel:", value)
			} else {
				fmt.Println("Not OK value from odd channel (closed)")
			}
		case value, ok := <-quitChannel:
			if ok {
				fmt.Println("from quit channel:", value)
			} else {
				fmt.Println("Not OK value from quit channel (closed)")
			}
			return
		}
	}
}
