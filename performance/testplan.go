package performance

import (
	. "tyrannosaurs/performance/component"
)

type TestPlan struct {
	Name        string
	ThreadGroup ThreadGroup
}

func (t *TestPlan) Start() {
}

func (t *TestPlan) initContext() {

}
