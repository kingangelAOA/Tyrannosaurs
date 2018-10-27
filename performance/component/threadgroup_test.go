package component

import (
	"testing"
)

func TestThreadGroup_start(t *testing.T) {
	tg := &ThreadGroup{
		Number: 20,
		Duration: 10,
	}
	flag1 := 0
	flag2 := 0
	tg.Start()
	println(flag1, flag2)
}