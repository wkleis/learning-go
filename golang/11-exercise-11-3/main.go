package main

// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type Users []user

func (s Users) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s Users) Less(i, j int) bool {
	if i == j {
		return false
	}
	pi := s[i]
	pj := s[j]
	cl := strings.Compare(pi.Last, pj.Last)
	if cl != 0 {
		return cl < 0
	}
	cf := strings.Compare(pi.First, pj.First)
	if cf != 0 {
		return cf < 0
	}
	return pi.Age < pj.Age
}

// Swap swaps the elements with indexes i and j.
func (s Users) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	u3 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is so good to see you",
			"Would you like me to take care of you, James",
			"I would really prefer to be a secret agent myself",
		},
	}
	u1 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James, you didn't",
			"Dear God, what has James done now?",
			"Can someone please tell me where james Bond is?",
		},
	}
	users := Users{u1, u2, u3}
	fmt.Println("--- not sorted ---")
	fmt.Println(users)
	for _, u := range users {
		sort.Strings(u.Sayings)
	}
	fmt.Println("--- sorted ---")
	sort.Sort(users)
	fmt.Println(users)

	fmt.Println("--- JSON ---")
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}
