package main

// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import (
	"fmt"
)

type t1 struct {
	name string
	size int
}

func (x *t1) f1() {
	fmt.Printf("%T ", x)
	fmt.Println("F1", x.name, x.size)
}
func (x t1) f2() {
	fmt.Printf("%T ", x)
	fmt.Println("F2", x.name, x.size)
}

type i1 interface {
	f1()
}

type i2 interface {
	f2()
}

func info1(x i1) {
	fmt.Print("info1 - ")
	x.f1()
}
func info2(x i2) {
	fmt.Print("info1 - ")
	x.f2()
}

func main() {
	t := t1{"Foo", 42}
	p := &t
	p.f1()
	t.f1() //this works because it is automatically translated to (&t1).f1(), see https://golang.org/ref/spec#Calls
	(t1{"Bar", 77}).f2()
	// (t1{"Bar", 77}).f1()   --> this would fail with ./prog.go:28:17: cannot call pointer method on T1 literal
	//see answer to https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver
	info2(t)
	info2(p) //the non-pointer traget type method is applicable to the pointer as well
	info1(p)
	// info1(t) --> would fail with:
	// cannot use t (type t1) as type i1 in argument to info1: t1 does not implement i1 (f1 method has pointer receiver)

}
