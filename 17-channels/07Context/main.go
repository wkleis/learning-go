// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	contextDemo2()
}
func contextDemo1() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("goroutines 1:", runtime.NumGoroutine())
	fmt.Println("context:", ctx)
	var ws sync.WaitGroup
	ws.Add(1)
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				ws.Done()
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("goroutines 2:", runtime.NumGoroutine())
	fmt.Println("context:", ctx)
	fmt.Println("About to cancel context")
	cancel()
	ws.Wait()
	fmt.Println("Cancelled context")
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("goroutines 3:", runtime.NumGoroutine())
}
func contextDemo2() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	outChannel := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done(): //context cancelled
				return // end the goroutine => not leak it
			case outChannel <- n: //out channel ready for writing:  write n
				n++
			}
		}
	}()
	return outChannel

}
