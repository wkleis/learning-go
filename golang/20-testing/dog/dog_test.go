package dog

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var dog = Dog{
	Name: "Fifi",
	Age:  2,
}

func TestDogYears(t *testing.T) {
	got := dog.DogYears()
	want := dog.Age * 7
	if got != want {
		t.Errorf("DogYears() = %d, want: %d (Age: %d)", got, want, dog.Age)
	}
	fmt.Println()
}

func TestDogInfo(t *testing.T) {
	got := dog.Info()
	// contains 3 sub tests
	t.Run("NotEmpty", func(t *testing.T) {
		if got == "" {
			t.Errorf("Info(): got empty string, want: non-empty string")
			t.FailNow()
		}
	})
	t.Run("ContainsName", func(t *testing.T) {
		want := dog.Name
		if !strings.Contains(got, want) {
			t.Errorf("Info(): got '%s', want: string that contains dog's name '%s'", got, want)
		}
	})
	t.Run("ContainsAge", func(t *testing.T) {
		want := strconv.Itoa(dog.Age)
		if !strings.Contains(got, want) {
			t.Errorf("Info(): got '%s', want: string that contains dog's age '%s'", got, want)
		}
	})
}

//run all benchmarks: go test --bench .

func BenchmarkDogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.DogYears()
	}
}
