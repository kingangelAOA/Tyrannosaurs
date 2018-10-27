package component

import (
	. "tyrannosaurs/constant"
	"context"
	"time"
	"sync"
)

var once sync.Once

type ThreadGroup struct {
	Name                    string `json:"Name"`
	Comments                string
	ActionAfterError        ACTION_AFTER_ERROR
	Number                  int
	RampUp                  int
	Duration                int64
	ConcurrentGoroutinesMap map[string]chan struct{}
	DoneMap                 map[string]chan bool
	Scenarios               []Scenario
}

func (t *ThreadGroup) addTask() {
	once.Do(func() {
		t.ConcurrentGoroutinesMap = map[string]chan struct{}{}
		t.DoneMap = map[string]chan bool{}
	})
	for _, s := range t.Scenarios {
		t.ConcurrentGoroutinesMap[s.Name] = make(chan struct{}, t.Number)
		t.DoneMap[s.Name] = make(chan bool)
		t.Scenarios = append(t.Scenarios, s)
		for i := 0; i < t.Number; i++ {
			t.ConcurrentGoroutinesMap[s.Name] <- struct{}{}
		}
	}
}

func (t *ThreadGroup) inputThread(ctx context.Context) {
	t.addTask()
	for k := range t.ConcurrentGoroutinesMap {
		go func(ctx context.Context, k string) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					<-t.DoneMap[k]
					t.ConcurrentGoroutinesMap[k] <- struct{}{}
				}
			}
		}(ctx, k)
	}
}

func (t *ThreadGroup) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	t.inputThread(ctx)
	for _, scenario := range t.Scenarios {
		go func(s Scenario) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					<-t.ConcurrentGoroutinesMap[s.Name]
					go func(s Scenario) {
						s.run()
						t.DoneMap[s.Name] <- true
					}(s)
				}
			}
		}(scenario)
	}
	t.stop(cancel)
}

func (t *ThreadGroup) stop(cancel func()) {
	time.Sleep(time.Duration(t.Duration) * time.Second)
	cancel()
}