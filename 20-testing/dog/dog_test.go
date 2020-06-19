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
	if got == "" {
		t.Errorf("Info(): got empty string, want: non-empty string")
		return
	}
	if !strings.Contains(got, dog.Name) {
		t.Errorf("Info(): got '%s', want: string that contains dog's name '%s'", got, dog.Name)
	}
	ageString := strconv.Itoa(dog.Age)
	if !strings.Contains(got, ageString) {
		t.Errorf("Info(): got '%s', want: string that contains dog's age '%s'", got, ageString)
	}
}

func ExampleDog_Info() {
	fmt.Println(dog.Info())
	//Output:
	//Dog 'Fifi' age: 2 (human years) or 14 (dog years)
}

//run all benchmarks: go test --bench .

func BenchmarkDogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.DogYears()
	}
}
