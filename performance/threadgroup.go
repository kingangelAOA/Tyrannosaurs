package performance

import (
	. "tyrannosaurs/constant"
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
