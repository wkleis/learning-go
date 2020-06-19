package sayings

import "fmt"

//Greet returns a string the greents the person with the specified name
func Greet(name string) string {
	return fmt.Sprintf("Welcome my dear %s.", name)
}
