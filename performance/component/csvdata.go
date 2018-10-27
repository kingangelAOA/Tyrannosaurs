package component

import (
	"os"
	"bufio"
	"io"
	"strings"
	"math/rand"
	. "tyrannosaurs/util"
	. "tyrannosaurs/constant"
)

type CSVData struct {
	Name      string
	FileName  string
	Params    string
	Separator string
	Loop      bool
	DataCache []map[string]string
	Index     int
	Base
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

func (c *CSVData) getTestData() (map[string]string, error) {
	if err := c.initCSVData(); err != nil {
		return nil, err
	}
	length := len(c.DataCache)
	if c.Loop {
		if length-1 < c.Index {
			c.Index = 0
		}
		result := c.DataCache[c.Index]
		c.Index ++
		return result, nil
	} else {
		return c.DataCache[rand.Intn(length)], nil
	}
}

func (c *CSVData) Notify() error {
	for _, o := range c.ObserverList {
		data, err := c.getTestData()
		if err != nil {
			return err
		}
		if err := o.Receive(data); err != nil {
			return err
		}
	}
	return nil
}
