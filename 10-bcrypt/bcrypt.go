package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	cryptBytes, error := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(s)
	fmt.Println(cryptBytes)
}
