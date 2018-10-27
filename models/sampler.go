package models

import "github.com/jinzhu/gorm"

type Sampler struct {
	gorm.Model
	ScenarioId   uint
	HttpRequest  HttpRequest
	Script       Script
	Consuming    int64
	TestDataList []TestData
}
