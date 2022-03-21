package convpack

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func OpenXlsx(path string) [][][]string {
	xlFile, err := xlsx.FileToSlice(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return xlFile
}
