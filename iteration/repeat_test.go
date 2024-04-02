package iteration

import (
	"fmt"
	"strings"
	"testing"
)

// func TestRepeat(t *testing.T) {
// 	repeated := Repeat("a")
// 	expected := "aaaaa"

// 	if repeated != expected {
// 		t.Errorf("expected %q but got %q", expected, repeated)
// 	}

// }

func TestRepeat(t *testing.T) {
	t.Run("repeat `n` number of times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("custom repeat matches standard lib repeat", func(t *testing.T) {
		c_repeat := Repeat("a", 10)
		sl_repeat := strings.Repeat("a", 10)

		if c_repeat != sl_repeat {
			t.Errorf("expected %q but got %q", sl_repeat, c_repeat)
		}
	})
}

// this is benchmarking and it follows exactly this structure
// where it must begin with "Benchmark" and take that parameter
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}
