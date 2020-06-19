package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Println(f)
	r := strings.NewReader("Just some text t be copied into the file")
	io.Copy(f, r)
}
