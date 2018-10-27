package component

import (
	. "tyrannosaurs/util"
)

type Base struct {
	ObserverList []Observer
}

func (b *Base) Attach(o Observer) {
	if !SliceExists(b.ObserverList, o) {
		b.ObserverList = append(b.ObserverList, o)
	}
}

func (b *Base) Detach(o Observer) {
	if SliceExists(b.ObserverList, o) {
		i := findObserver(b.ObserverList, o)
		if i != -1 {
			b.ObserverList = append(b.ObserverList[:i], b.ObserverList[i+1:]...)
		}
	}
}

func findObserver(observers []Observer, observer Observer) int {
	for i, o := range observers {
		if o == observer {
			return i
		}
	}
	return -1
}