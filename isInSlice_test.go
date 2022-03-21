package convpack

import (
	"testing"
)

func TestIsInSlice(t *testing.T) {
	expected := 3
	actual := 0

	str := "a"
	strSlice := []string{"a", "b", "c", "d"}
	num := 3
	intSlice := []int{1, 2, 3, 4, 5}
	ru := 'r'
	runeSlice := []rune{'q', 'w', 'e', 'r'}

	if IsInSlice(str, strSlice) {
		actual++
	}

	if IsInSlice(num, intSlice) {
		actual++
	}

	if IsInSlice(ru, runeSlice) {
		actual++
	}

	if expected != actual {
		t.Fatal("Fatal! expected : ", expected, " actual : ", actual)
	}
}
