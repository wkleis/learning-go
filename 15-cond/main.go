package main
// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	data := []int{}
	var mutex sync.Mutex
	var lock sync.Mutex
	var cond = sync.NewCond(&lock)
	var wg sync.WaitGroup
	gs := 4
	counter := 0
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			counter++
			me := counter
			mutex.Unlock()
			fmt.Println("Goroutine", me, "started")
			cond.L.Lock()
			fmt.Println("Goroutine", me, "locked")
			for len(data) < 4 {
				cond.Wait()
			}
			fmt.Println("Goroutine", me, "before: ", data)
			for k := 0; k < 4; k++ {
				fmt.Println("  Goroutine", me, "data:", data[k])
			}
			data = data[4:]
			fmt.Println("Goroutine", me, "after:", data)
			cond.L.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("routines started")
	time.Sleep(time.Second)
	j := 0
	for runtime.NumGoroutine() > 1 {
		cond.L.Lock()
		time.Sleep(10 * time.Millisecond)
		data = append(data, j*2)
		data = append(data, j*2+1)
		j++
		cond.Broadcast()
		cond.L.Unlock()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
