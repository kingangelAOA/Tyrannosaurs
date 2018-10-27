package performance

import (
	"testing"
	"io/ioutil"
	"encoding/json"
)

func TestTestPlan_Start(t *testing.T) {
	plan, _ := ioutil.ReadFile("/Users/kingangel/work/go/src/tyrannosaurs/example.json")
	var testplans []TestPlan
	if err := json.Unmarshal(plan, &testplans); err != nil {
		println(err.Error())
	}
	for _, t := range testplans {
		t.ThreadGroup.Start()
	}
}
