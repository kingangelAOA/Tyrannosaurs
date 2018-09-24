package models

import (
	"os"
	. "github.com/kingangelAOA/tyrannosaurs/util"
	. "github.com/kingangelAOA/tyrannosaurs/constant"
	"bufio"
	"io"
	"strings"
	"sync"
	"math/rand"
)

var once sync.Once

type CSVData struct {
	Name      string
	FileName  string
	Params    string
	Separator string
	Loop      bool
	DataCache []map[string]string
	index     int
}

func (c *CSVData) initCSVData() (e error) {
	once.Do(func() {
		f, err := os.Open(c.FileName)
		defer f.Close()
		if err != nil {
			Log.Fatal(OpenTestFileError.Error())
			e = OpenTestFileError
		}
		if c.Params == "" {
			Log.Fatal(CVSDataParamError.Error())
			e = CVSDataParamError
		}
		bfRd := bufio.NewReader(f)
		keys := strings.Split(c.Params, c.Separator)
		var resultList []map[string]string
		for {
			resultMap := make(map[string]string)
			line, err := bfRd.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			values := strings.Split(string(line), c.Separator)
			for i, v := range values {
				if i < len(keys) {
					resultMap[keys[i]] = v
				}
			}
			resultList = append(resultList, resultMap)
		}
		c.DataCache = resultList
	})
	return e
}

func (c *CSVData) GetTestData() (map[string]string, error) {
	if err := c.initCSVData(); err != nil {
		return nil, err
	}
	length := len(c.DataCache)
	if c.Loop {
		if length - 1 < c.index {
			c.index = 0
		}
		result := c.DataCache[c.index]
		c.index ++
		return result, nil
	} else {
		return c.DataCache[rand.Intn(length)], nil
	}
}
