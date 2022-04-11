package convpack

import (
	"testing"
)

func TestDeleteDup(t *testing.T) {
	expected := 3
	actual := 0

	sl := []string{"A", "B", "C", "A", "B"}
	il := []int{1, 2, 3, 4, 1, 2, 3, 5, 6, 7, 5}
	bl := []bool{true, true, true, true, false}

	if len(DeleteDup(sl)) == 3 {
		actual++
	}
	if len(DeleteDup(il)) == 7 {
		actual++
	}
	if len(DeleteDup(bl)) == 2 {
		actual++
	}

	if expected != actual {
		t.Fatal("Fatal! expected : ", expected, " actual : ", actual)
	}
}
