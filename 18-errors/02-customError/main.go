package main

import (
	"fmt"
)

type FooError struct {
	message string
	code    int
}

func (e *FooError) Error() string {
	return fmt.Sprintf("Foo Error %s (code: %d)", e.message, e.code)
}

func main() {
	var err error
	func() {
		var v int
		v, err = foo(2)
		fmt.Println(v)
		v, err = foo(-1)
		fmt.Println(v)

	}()
	if err != nil {
		fmt.Println(err)
	}
}
func foo(n int) (int, error) {
	if n == -1 {
		return 0, &FooError{message: "Foo: Illegal value (-1)", code: 33}
	}
	return n, nil
}
