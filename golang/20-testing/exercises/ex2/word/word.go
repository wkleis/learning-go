//from https://github.com/GoesToEleven/go-programming/blob/master/code_samples/010-ninja-level-thirteen/02/01-code-starting/word/main.go

package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

//UseCount returns all the "words" in strings with the number of their ocurrences in the string
//returns a map where the "word" is key and value in the count
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count the "words" in the string s
func Count(s string) int {
	return len(strings.Fields(s))
}
