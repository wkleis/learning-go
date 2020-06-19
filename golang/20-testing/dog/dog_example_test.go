package dog

import (
	"fmt"
)

func ExampleDog_Info() {
	var dog = Dog{
		Name: "Fifi",
		Age:  2,
	}
	fmt.Println(dog.Info())
	//Output:
	//Dog 'Fifi' age: 2 (human years) or 14 (dog years)
}
