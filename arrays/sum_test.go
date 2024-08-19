package arrays

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: 6
}

func TestSumAll(t *testing.T) {
	t.Run("Two collections of any size", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{2, 3, 4})
		want := []int{6, 9}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(sums)
	// Output: [6 9]
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Sum of tails of given collections", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		want := []int{5, 7}

		checkSums(t, got, want)

	})

	t.Run("Safely sum an empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})

}

func ExampleSumAllTails() {
	sumOfTails := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(sumOfTails)
	// Output: [5 7]
}
