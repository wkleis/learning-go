package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password123"
	hashedPw, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(password)
	fmt.Println(hashedPw)
	loginPassword := "password123"
	error = bcrypt.CompareHashAndPassword(hashedPw, []byte(loginPassword))
	fmt.Println(loginPassword)
	if error == nil {
		fmt.Println("login OK")
	} else {
		fmt.Println("login failed")
	}

}
