package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of varying size", func(t *testing.T) {
		mySlice := []int{1,2,3}

		got := Sum(mySlice)
		want := 6

		if got !=  want {
			t.Errorf("got %d want %d, received %v", got, want, mySlice)
		}
	})
}
