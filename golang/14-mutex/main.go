package main
// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("NumGoRoutine", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("NumGoRoutine", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}
