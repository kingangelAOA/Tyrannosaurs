package component

import (
	. "tyrannosaurs/constant"
	"github.com/tidwall/gjson"
	"fmt"
)

const ErrorFormat = "json: %s, path: %s, expect: %s, actual: %s"

type JsonAssert struct {
	Path   string
	Expect string
	BaseAssert
}

func (j *JsonAssert) Receive(message interface{}) error {
	if m, ok := message.(string); ok {
		if !gjson.Valid(m) {
			return JsonAssertFormatError
		}
		actual := gjson.Get(m, j.Path).String()
		ba := BaseAssert{}
		if actual == j.Expect {
			ba.IsTrue = true
			ba.ErrMessage = ""
		} else {
			ba.IsTrue = false
			ba.ErrMessage = fmt.Sprintf(ErrorFormat, m, j.Path, j.Expect, actual)
		}
		return nil
	}
	return JsonAssertReceiveIsNotString
}
