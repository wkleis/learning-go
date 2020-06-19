package sayings

import (
	"fmt"
	"testing"
)

func ExampleGreet() {
	fmt.Println(Greet("James"))
	//Output:
	//Welcome my dear James.
}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}

/*
Sample run:
>go test -bench Greet
goos: linux
goarch: amd64
pkg: github.com/wkleis/learning-go/20-testing/sayings
BenchmarkGreet-4   	 7494589	       152 ns/op
PASS
ok  	github.com/wkleis/learning-go/20-testing/sayings	1.304s
*/
