// from Udemy course https://www.udemy.com/course/learn-how-to-code/ by Todd McLeod
package word

import (
	"strings"
	"testing"

	"github.com/wkleis/learning-go/golang/20-testing/exercises/ex2/quote"
)

func TestCount(t *testing.T) {
	in := quote.SunAlso
	want := 1349
	got := Count(in)
	if want != got {
		t.Errorf("Count: want %d, got %d", want, got)
	}

}
func TestUseCount(t *testing.T) {
	in := "Alpha Beta Gamma \t Alpha Alpha \t "
	for i := 0; i < 10; i++ {
		in = strings.Join([]string{in, "Gamma"}, " ")
	}

	for i := 0; i < 100; i++ {
		in = strings.Join([]string{in, "Delta"}, " ")
	}

	want := map[string]int{
		"Alpha": 3,
		"Beta":  1,
		"Gamma": 11,
		"Delta": 100,
	}

	got := UseCount(in)
	if len(want) != len(got) {
		t.Fatalf("UseCount: wanted map with %d elements, got one with %d (%v)", len(want), len(got), got)
	}

	for key, wantValue := range want {
		gotValue, ok := got[key]
		if !ok {
			t.Errorf("UseCount: expected word %s missing in map %v.", key, got)
			continue
		}
		if wantValue != gotValue {
			t.Errorf("UseCount: count for word %s: want %d, got %d.", key, wantValue, gotValue)
		}
	}

}

func BenchmarkCount(b *testing.B) {
	q := quote.SunAlso
	for i := 0; i < b.N; i++ {
		Count(q)
	}
}

func BenchmarkUseCount(b *testing.B) {
	q := quote.SunAlso
	for i := 0; i < b.N; i++ {
		UseCount(q)
	}
}
