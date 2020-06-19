// from https://github.com/GoesToEleven/go-programming/blob/master/code_samples/010-ninja-level-thirteen/02/01-code-starting/main.go

package main

import (
	"fmt"

	"github.com/wkleis/learning-go/golang/20-testing/exercises/ex2/quote"
	"github.com/wkleis/learning-go/golang/20-testing/exercises/ex2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println("--------------------------")
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
