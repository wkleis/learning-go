package main

import (
	"fmt"
)

func main() {
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
}
