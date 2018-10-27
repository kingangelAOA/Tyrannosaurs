package component

import (
	"testing"
)

type Test struct {
	data map[string]string
}

func replace(m map[string]string)  {
	m["a"] = "acadfasdfasdfas"
}

func TestJsonExtractor_Receive(t *testing.T) {
	test := &Test{
		data: map[string]string{"a" : "1", "b": "c"},
	}

	replace(test.data)
	println(test.data["a"])
}