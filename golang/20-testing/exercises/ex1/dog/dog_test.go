package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	humanYears := 3
	want := 21
	t.Run("Years", func(t *testing.T) {
		got := Years(humanYears)
		if want != got {
			t.Errorf("Years(%d): want %d, got: %d", humanYears, got, want)
		}
	})
	t.Run("YearsTwo", func(t *testing.T) {
		got := YearsTwo(humanYears)
		if want != got {
			t.Errorf("YearsTwo(%d): want %d, got: %d", humanYears, got, want)
		}
	})
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}

/*
benchmarking results:
go test --bench . --benchtime 100000x
goos: linux
goarch: amd64
pkg: github.com/wkleis/learning-go/20-testing/exercises/ex1/dog
BenchmarkYears-4      	  100000	         0.373 ns/op
BenchmarkYearsTwo-4   	  100000	     15403 ns/op
PASS

*/

func ExampleYears() {
	fmt.Println(Years(2))
	//Output:
	//14
}
