package component

import (
	"testing"
	"fmt"
	"container/list"
)

func TestGetTestData(t *testing.T) {
	csvData := CSVData{
		FileName:  "/Users/kingangel/work/go/src/tyrannosaurs/info.log",
		Separator: ";",
		Params:    "a",
		Loop:      false,
	}
	for {
		result, err := csvData.getTestData()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

func TestCSVData_Attach(t *testing.T) {
	l := list.New()
	l.PushBack(0)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
}
