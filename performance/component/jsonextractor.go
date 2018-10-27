package component

import (
	. "tyrannosaurs/constant"
	"github.com/tidwall/gjson"
	"fmt"
	"strings"
)

type JsonExtractor struct {
	Name       string
	VarValue   map[string]string
	VarExpList []VarExp
	Base
}

type VarExp struct {
	Key          string
	Path         string
	DefaultValue string
}

func (je *JsonExtractor) newExtractor(j string) error {
	for _, varExp := range je.VarExpList {
		if !gjson.Valid(j) {
			return JsonExtractorJsonFormatError
		}
		result := gjson.Get(j, strings.Replace(varExp.Path, "$.", "", -1))
		if !result.Exists() {
			return fmt.Errorf("json Path '%s' did not found ", varExp.Path)
		}
		je.VarValue[varExp.Key] = result.String()
	}
	return nil
}

func (je *JsonExtractor) Receive(message interface{}) error {
	if m, ok := message.(string); !ok {
		je.newExtractor(m)
		return nil
	}
	return JsonExtractorReceiveFormatError
}

func (je *JsonExtractor) Notify() error {
	for _, o := range je.ObserverList {
		if err := o.Receive(je.VarValue); err != nil {
			return err
		}
	}
	return nil
}
