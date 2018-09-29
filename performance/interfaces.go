package performance

type Sampler interface {
	Fetch() (string, error)
}

type Assertions interface {
	GetResult() string
}

type ConfigElement interface {

}

type PreProcessors interface {

}

type PostProcessors interface {

}