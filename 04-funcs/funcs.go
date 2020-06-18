package main

import (
	"fmt"
)

// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

type context struct {
	size   int
	number int
	name   string
}

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
	ctx := context{
		size:   2,
		number: 11,
		name:   "foo",
	}
	f2 := foo2(ctx)
	fmt.Println(f2())
	bar("Hurz",
		func(s string) {
			fmt.Println("callback called with: ", s)
		})
	closureDemo()

}

//function that returns function with string variable bound to varibale from function closure
func foo(s string) func() string {
	return func() string {
		return fmt.Sprintf("The word is: %v", s)
	}
}

//function that returns functions with struct variable bound to variable from function closure
func foo2(c context) func() string {
	return func() string {
		return fmt.Sprintf("size: %d, number: %d, %s", c.size, c.number, c.name)
	}
}

//function with callback
func bar(s string, cb func(s string)) {
	x := fmt.Sprintf("The message is '%s'!!!", s)
	cb(x)
}
func closureDemo() {
	i1 := incrementor()
	i2 := incrementor()
	fmt.Println(i1())
	fmt.Println(i1())
	fmt.Println(i1())
	fmt.Println(i1())
	fmt.Println(i2())
	fmt.Println(i2())
	fmt.Println(i2())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
