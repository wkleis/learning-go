package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	raceConditionDemo()
}
func raceConditionDemo() {
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("NumGoRoutine", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("NumGoRoutine", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter", counter)
}
