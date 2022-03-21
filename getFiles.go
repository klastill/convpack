package convpack

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
)

func GetFiles(directory string, s ...string) []fs.FileInfo {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(s) == 0 {
		return files
	}

	var useFileList []fs.FileInfo
	var extensions []string

	extensions = append(extensions, s...)

	for _, file := range files {
		ext := string(file.Name()[strings.LastIndex(file.Name(), ".")+1:])
		if IsInSlice(ext, extensions) {
			useFileList = append(useFileList, file)
		}
	}

	return useFileList
}
