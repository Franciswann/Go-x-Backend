package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
		// Repeat("a")
	}
}

func ExampleRepeat() {
	result := Repeat("R", 3)
	fmt.Println(result)
	// Output: RRR
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := ("aaaa")

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
