//This is an exercise from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

//Package dog contains useless functions about the world of dogs
package dog

import "fmt"

//Dog is a datatype with information about one dog
type Dog struct {
	Name string
	Age  int
}

//Info returns  astring with informaion about the dog
func (d Dog) Info() string {
	return fmt.Sprintf("Dog '%s' age: %d (human years) or %d (dog years)", d.Name, d.Age, d.DogYears())
}

//DogYears returns the age of the dog as dog ywears (not human years)
func (d Dog) DogYears() int {
	return d.Age * 7
}
