package main

//slight variation of exmaple from: https://blog.golang.org/defer-panic-and-recover
import (
	"fmt"
)

func main() {
	fmt.Println("main: start")
	a()
	fmt.Println("main: end")
}

func a() {
	fmt.Println("a: start")
	fmt.Println("a: deferred")
	b()
	fmt.Println("a: end")

}

type MySpecialError struct {
	error
	Code int
}

func b() {
	fmt.Println("b: start")
	defer func() {
		fmt.Println("b: deferred")
		if r := recover(); r != nil {
			fmt.Printf("b: Recovered from '%v'\n", r)
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println("b: i=", i)
		c(i)
	}
	fmt.Println("b: end")

}

func c(p int) {
	fmt.Println("c: start")
	defer fmt.Println("c: deferred")
	d(p)
	fmt.Println("c: end")

}
func d(p int) {
	fmt.Println("d: start")
	defer fmt.Println("d: deferred")
	if p >= 2 {
		fmt.Printf("d: about to panic now")
		panic(fmt.Errorf("Panic in func d!. (p=%d)", p))
	}
	fmt.Println("d: end")
}
