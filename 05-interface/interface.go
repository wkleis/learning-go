package main

// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import (
	"fmt"
	"math"
)

type square struct {
	a float64
}
type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.a * s.a
}
func (c circle) area() float64 {
	return c.r * c.r * math.Pi
}

type shape interface {
	area() float64
}

func main() {
	squares := []square{{5.5}, {2.0}, {3.99}}
	circles := []circle{{1.0}, {2.0}, {3}, {77.24}}
	n := len(squares) + len(circles)
	pos := 0
	shapes := make([]shape, n, n)
	fmt.Println("squares")
	for _, x := range squares {
		fmt.Printf("%v\t%v\n", x, x.area())
		shapes[pos] = x
		pos++
	}
	fmt.Println("Circles")
	for _, x := range circles {
		fmt.Printf("%v\t%v\n", x, x.area())
		shapes[pos] = x
		pos++
	}

	fmt.Println("Shapes")
	for _, x := range shapes {
		info(x)
	}
}

func info(s shape) {
	fmt.Printf("%T\t%v\t%v\n", s, s, s.area())
}
