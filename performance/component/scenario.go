package component

type Scenario struct {
	Name          string
	Comment       string
	Throughput    float32
	JsonAssert    []JsonAssert
	CSVDataList   []CSVData
	JsonExtractor JsonExtractor
	UserParams    UserParams
	Http          []Http
}

func (s *Scenario) run() {
	if len(s.CSVDataList) > 0 {
		for _, cdl := range s.CSVDataList {
			cdl.Attach()
		}
	}
}
