package postprocessors

import (
	"encoding/json"
	"fmt"
	"k8s.io/client-go/util/jsonpath"
	"bytes"
	. "tyrannosaurs/constant"
)

type JsonExtractor struct {
	Name       string
	VarValue   map[string]string
	VarExpList []VarExp
}

type VarExp struct {
	Var     string
	path    string
	Default string
}

func (je *JsonExtractor) NewExtractor(j string) error {
	for _, varExp := range je.VarExpList {
		input := []byte(j)
		var data interface{}
		if err := json.Unmarshal(input, &data); err != nil {
			return JsonExtractorJsonFormatError
		}
		jp := jsonpath.New(je.Name)
		jp.AllowMissingKeys(false)
		if err := jp.Parse(varExp.Var); err != nil {
			return JsonExtractorPathFormatError
		}
		buf := new(bytes.Buffer)
		if err := jp.Execute(buf, data); err != nil {
			return fmt.Errorf("json path '%s' did not found ", varExp.path)
		}
		je.VarValue[varExp.Var] = buf.String()
	}
	return nil
}
