package convpack

import (
	"testing"
)

func TestFilter(t *testing.T) {
	expected := 5
	actual := 0

	l1 := []string{"1", "2", "3", "4", "5", "6"}
	actual += len(Filter(l1, func(l string) bool {
		return l == "1"
	}))

	l2 := []int{1, 2, 3, 4, 5, 6, 7}
	actual += len(Filter(l2, func(l int) bool {
		return l > 3
	}))

	if expected != actual {
		t.Fatal("Fatal! expected : ", expected, " actual : ", actual)
	}
}
