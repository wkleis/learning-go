package main

import (
	"fmt"
)

type T1 struct {
	name string
	size int
}

func (x *T1) f1() {
	fmt.Printf("%T ", x)
	fmt.Println("F1", x.name, x.size)
}
func (x T1) f2() {
	fmt.Printf("%T ", x)
	fmt.Println("F2", x.name, x.size)
}

func main() {
	t := T1{"Foo", 42}
	p := &t
	p.f1()
	t.f1() //this works because it is automatically translated to (&t1).f1(), see https://golang.org/ref/spec#Calls
	(T1{"Bar", 77}).f2()
	// (T1{"Bar", 77}).f1()   --> this would fail with ./prog.go:28:17: cannot call pointer method on T1 literal
	//see answer to https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver

}
