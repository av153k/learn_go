package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Add two given numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("sum %d expected %d", sum, expected)
		}
	})
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
