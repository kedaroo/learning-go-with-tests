package arraysandslices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of varying size", func(t *testing.T) {
		mySlice := []int{1, 2, 3}

		got := Sum(mySlice)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, received %v", got, want, mySlice)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	got := SumAll(slice1, slice2)
	want := []int{6, 15}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got %v want %v, input %v, %v", got, want, slice1, slice2)
		}
	}
}
