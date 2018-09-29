package performance

type Sampler interface {
	Fetch() ([]byte, error)
}

type Assertions interface {
	GetResult() []byte
	IsFailed() bool
}

type ConfigElement interface {
	GetValueByKey(key string) string
}

type PreProcessors interface {
	Process()
}

type PostProcessors interface {
	Process()
}