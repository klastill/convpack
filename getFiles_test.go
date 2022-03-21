package convpack

import (
	"fmt"
	"os"
	"testing"
)

func TestGetFIlesFiltering(t *testing.T) {
	expected := 5
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	files := GetFiles(dir, "go")

	actual := len(files)

	if expected != actual {
		t.Fatal("Fatal! expected : ", expected, " actual : ", actual)
	}
}
