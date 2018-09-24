package models

type TestPlan struct {
	Name         string
	Params       map[string]string
	ThreadGroups *[]ThreadGroup
}

type Param struct {
	Key   string
	Value string
}

