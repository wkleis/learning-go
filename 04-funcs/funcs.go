package main

import (
	"fmt"
)

func main() {
	//function expression
	funcs := []func(string){
		func(s string) {
			fmt.Println("foo:", s)
		},
		func(s string) {
			fmt.Println("bar:", s)
		},
		func(s string) {
			fmt.Println("blub:", s)
		},
	}
	for _, f := range funcs {
		fmt.Printf("%T\n", f)
		f("Hurz")
	}
	fmt.Println("-------------")
	//function as return value
	f := foo("Hurz")
	fmt.Println(f())
}

func foo(s string) func() string {
	return func() string {
		return fmt.Sprintf("The word is: %v", s)
	}
}
