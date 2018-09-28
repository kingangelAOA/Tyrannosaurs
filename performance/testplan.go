package performance

type TestPlan struct {
	Name         string
	Params       map[string]string
	ThreadGroups *[]ThreadGroup
}

