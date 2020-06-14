package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      35,
		"Moneypenny": 25,
		"Q":          55,
	}

	printFromMap(m, "James")
	printFromMap(m, "Tom")
	m["Freddy"] = 44
	fmt.Println("-------")
	for key, val := range m {
		fmt.Println(key, val)
	}
	del(m, "Freddy")
	fmt.Println(m)

}

func printFromMap(m map[string]int, n string) {
	if v, ok := m[n]; ok {
		fmt.Println(n, v)
	} else {
		fmt.Println(n, "does not exist")
	}

}
func del(m map[string]int, n string) {
	if _, ok := m[n]; ok {
		fmt.Println("deleting", n)
		delete(m, n)
	} else {
		fmt.Println(n, "does not exist")
	}

}
