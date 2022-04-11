package convpack

import (
	"testing"
)

func TestIsInSlice(t *testing.T) {
	expected := 3
	actual := 0

	strSlice := []string{"1", "2", "3", "4"}
	intSlice := []int{1, 2, 3, 4, 5}
	runeSlice := []rune{'1', '2', '3', '4'}

	if IsInSlice("1", strSlice) {
		actual++
	}

	if IsInSlice(3, intSlice) {
		actual++
	}

	if IsInSlice('2', runeSlice) {
		actual++
	}
	if IsInSlice(2, runeSlice) {
		actual++
	}
	if IsInSlice('2', intSlice) {
		actual++
	}
	if expected != actual {
		t.Fatal("Fatal! expected : ", expected, " actual : ", actual)
	}
}
