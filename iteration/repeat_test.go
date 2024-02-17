package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 10)
	}
}

func BenchmarkStandardRepeat(b *testing.B) {
	for range b.N {
		strings.Repeat("a", 10)
	}
}

// Results of benchmarks of our custom repeat vs strings.Repeat
// BenchmarkRepeat-8                6028188               175.1 ns/op
// BenchmarkStandardRepeat-8       36641360                31.39 ns/op

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}