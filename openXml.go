package convpack

import (
	"fmt"

	"github.com/beevik/etree"
)

func OpenXml(path string) *etree.Element {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		fmt.Println(err.Error())
	}
	rows := doc.SelectElement("rows")
	return rows
}
