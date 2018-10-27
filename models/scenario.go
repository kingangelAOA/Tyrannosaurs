package models

import "github.com/jinzhu/gorm"

type Scenario struct {
	gorm.Model
	ThreadId   uint
	Comment    string
	Throughput float32
	Samplers  []Sampler
}

