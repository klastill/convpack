package convpack

import (
	"testing"
)

func TestSet(t *testing.T) {
	expected := 3
	actual := 0

	s1 := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}
	s2 := []string{
		"a",
		"b",
		"c",
		"f",
		"g",
	}
	s3 := []string{
		"a",
		"b",
		"c",
		"d",
		"f",
	}
	s4 := []string{
		"a",
		"b",
		"c",
		"d",
		"g",
	}
	its_ans := []string{"a", "b", "c"}
	uni_ans := []string{"a", "b", "c", "d", "e", "f", "g"}
	diff_ans := []string{"d", "e"}

	its := GetIntersection(s1, s2, s3, s4)
	uni := GetUnion(s1, s2, s3, s4)
	diff := GetDefference(s1, its_ans)

	if len(its) == len(its_ans) {
		valid := true
		for _, item := range its {
			if !IsInSlice(item, its_ans) {
				valid = false
				break
			}
		}
		if valid {
			actual++
		}
	}
	if len(uni) == len(uni_ans) {
		valid := true
		for _, item := range uni {
			if !IsInSlice(item, uni_ans) {
				valid = false
				break
			}
		}
		if valid {
			actual++
		}
	}
	if len(diff) == len(diff_ans) {
		valid := true
		for _, item := range diff {
			if !IsInSlice(item, diff_ans) {
				valid = false
				break
			}
		}
		if valid {
			actual++
		}
	}
	if expected != actual {
		t.Fatal("Fatal! expected : ", expected, " actual : ", actual)
	}
}
