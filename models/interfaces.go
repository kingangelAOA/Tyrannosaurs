package models

type Sampler interface {
	Fetch() (string, error)
}

type Extractor interface {
	GetValue() string
}
