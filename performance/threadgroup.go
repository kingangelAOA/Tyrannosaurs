package performance

import (
	. "github.com/kingangelAOA/tyrannosaurs/constant"
)

type ThreadGroup struct {
	Name             string
	Comments         string
	ActionAfterError ACTION_AFTER_ERROR
	Number           int
	RampUp           int
	LoopCount        int
	Duration         int
	Tasks            []func()
}
