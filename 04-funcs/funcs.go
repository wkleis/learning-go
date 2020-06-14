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
	bar("Hurz",
		func(s string) {
			fmt.Println("callback called with: ", s)
		})
}

func foo(s string) func() string {
	return func() string {
		return fmt.Sprintf("The word is: %v", s)
	}
}

//function with callback
func bar(s string, cb func(s string)) {
	x := fmt.Sprintf("The message is '%s'!!!", s)
	cb(x)
}
