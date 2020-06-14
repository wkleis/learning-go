package main

import (
	"fmt"
	"sort"
	"strings"
)

type person struct {
	first string
	last  string
}

type Persons []person

func (s Persons) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s Persons) Less(i, j int) bool {
	if i == j {
		return false
	}
	pi := s[i]
	pj := s[j]
	cl := strings.Compare(pi.last, pj.last)
	if cl != 0 {
		return cl < 0
	}
	return strings.Compare(pi.first, pj.first) < 0

}

// Swap swaps the elements with indexes i and j.
func (s Persons) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	people := Persons{
		{first: "James", last: "Bond"},
		{first: "Frank", last: "Zappa"},
		{first: "James", last: "Brown"},
		{first: "Alan", last: "Turing"},
		{first: "Richard", last: "Strauss"},
		{first: "Johann", last: "Strauss"},
		{first: "Albert", last: "Einstein"},
	}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)

}
