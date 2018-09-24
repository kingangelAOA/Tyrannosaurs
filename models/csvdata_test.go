package models

import (
	"testing"
	"fmt"
)

func TestGetTestData(t *testing.T) {
	csvData := CSVData{
		FileName:  "/Users/kingangelTOT/go_work/src/github.com/kingangelAOA/tyrannosaurs/info.log",
		Separator: ";",
		Params:    "a",
		Loop:      false,
	}
	for {
		result, err := csvData.GetTestData()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	}
}
